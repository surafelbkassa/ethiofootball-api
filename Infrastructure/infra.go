package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	"github.com/redis/go-redis/v9"
)

// RedisConnect creates and returns a redis client using env vars
func RedisConnect() *redis.Client {
	ctx := context.Background()
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Username: redisUsername,
		Password: redisPassword,
		DB:       0,
	})

	// smoke-set (non-fatal)
	err := rdb.Set(ctx, "ethiofb:ping", "pong", 10*time.Second).Err()

	if err != nil {
		panic(err)
	}
	
	return rdb
}

// FetchFixturesFromAPI calls API-Football and returns simplified maps for Repository layer.
// Keys returned in each map:
//
//	id, league, home_id, away_id, date, status, score, home_logo, away_logo, last_update
//
// Accepts:
//
//	league: either "EPL" (will map to 39) or a numeric league id string (e.g. "39")
//	team: optional team id (numeric string) — names are NOT searched here
//	from,to: optional dates in YYYY-MM-DD
//
// Returns error if API key missing or upstream error.
func FetchFixturesFromAPI(league, team, season, from, to string) ([]map[string]string, error) {
	apiKey := os.Getenv("API_FOOTBALL_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("missing API_FOOTBALL_KEY in .env")
	}

	// Resolve league param: allow "EPL" -> 39, numeric IDs as passed
	leagueID := 0
	if league == "EPL" {
		leagueID = 39
	} else {
		// try numeric
		if n, err := strconv.Atoi(league); err == nil {
			leagueID = n
		} else {
			return nil, fmt.Errorf("unknown league code: %s (use 'EPL' or numeric league id)", league)
		}
	}

	base := "https://v3.football.api-sports.io"
	endpoint := "/fixtures"
	params := url.Values{}
	params.Set("league", strconv.Itoa(leagueID))
	// season optional; API often needs season for historical queries; leaving unset uses API default/current
	if from != "" {
		params.Set("from", from) // YYYY-MM-DD
	}
	if to != "" {
		params.Set("to", to)
	}
	if season != "" {
		params.Set("season", season)
	}

	if team != "" {
		// allow numeric team id only
		if _, err := strconv.Atoi(team); err == nil {
			params.Set("team", team)
		}
	}

	u := fmt.Sprintf("%s%s?%s", base, endpoint, params.Encode())

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-apisports-key", apiKey)
	// optional: req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 12 * time.Second}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// propagate upstream errors to caller
	if res.StatusCode >= 400 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("api error %d: %s", res.StatusCode, string(b))
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Use your domain.APIResponse shapes (you already defined them in Domain)
	var apiResp domain.APIResponse
	if err := json.Unmarshal(b, &apiResp); err != nil {
		// parsing failed - return raw error to help debugging
		return nil, fmt.Errorf("failed to unmarshal API response: %w", err)
	}

	out := make([]map[string]string, 0, len(apiResp.Response))
	now := time.Now().UTC().Format(time.RFC3339)

	for _, r := range apiResp.Response {
		idStr := strconv.Itoa(r.Fixture.ID)
		date := r.Fixture.Date // API returns ISO time string (RFC3339)
		homeName := r.Teams.Home.Name
		awayName := r.Teams.Away.Name
		homeLogo := r.Teams.Home.Logo
		awayLogo := r.Teams.Away.Logo

		// score/status handling
		scoreStr := ""
		status := "SCHEDULED"
		if r.Goals.Home != nil && r.Goals.Away != nil {
			scoreStr = fmt.Sprintf("%d-%d", *r.Goals.Home, *r.Goals.Away)
			status = "FT"
		} else if r.Fixture.Date != "" {
			// if fixture date in future, keep SCHEDULED; otherwise leave status as provided, if any
			status = "SCHEDULED"
		}

		m := map[string]string{
			"id":          idStr,
			"league":      r.League.Name,
			"home_id":     homeName,
			"away_id":     awayName,
			"date":        date,
			"status":      status,
			"score":       scoreStr,
			"home_logo":   homeLogo,
			"away_logo":   awayLogo,
			"last_update": now,
		}
		out = append(out, m)
	}

	return out, nil
}
