package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	"github.com/redis/go-redis/v9"
)

type StandingsRepo struct {
	apiURL string
	rdb    *redis.Client
}

func NewStandingsRepo(rdb *redis.Client) domain.IStandingsRepo {
	return &StandingsRepo{
		apiURL: "https://v3.football.api-sports.io/standings",
		rdb:    rdb,
	}
}


func (r *StandingsRepo) GetStandings(ctx context.Context, leagueID, season int) (*domain.StandingsResponse, error) {

    key := fmt.Sprintf("st:%d:%d", leagueID, season)
    // Try cache first
    cachedData, err := r.rdb.Get(ctx, key).Bytes()
    if err == nil {
        var standingsResponse domain.StandingsResponse
        if err := json.Unmarshal(cachedData, &standingsResponse); err == nil {
            return &standingsResponse, nil
        }
    } else if err != redis.Nil {
        fmt.Printf("Redis error: %v\n", err)
    }

    // Fetch from API
    API_KEY := os.Getenv("API_SPORTS_API_KEY")
    url := fmt.Sprintf("%s?league=%d&season=%d", r.apiURL, leagueID, season)

    client := &http.Client{}
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return nil, domain.ErrInternalServer
    }

    req.Header.Set("x-rapidapi-key", API_KEY)
    req.Header.Set("x-rapidapi-host", "v3.football.api-sports.io")

    res, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        return nil, domain.ErrInternalServer
    }
    defer res.Body.Close()

    body, err := io.ReadAll(res.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return nil, domain.ErrInternalServer
    }

    var apiResponse domain.StandingAPIResponse
    if err := json.Unmarshal(body, &apiResponse); err != nil {
        return nil, domain.ErrInternalServer
    }

    if len(apiResponse.Response) == 0 {
        return nil, fmt.Errorf("no standings found")
    }

    league := apiResponse.Response[0].League
    simplified := &domain.StandingsResponse{
        LeagueID:   league.ID,
        LeagueName: league.Name,
        Country:    league.Country,
        Season:     league.Season,
    }

    for _, group := range league.Standings {
        var groupStandings []domain.Standing
        for _, t := range group {
            groupStandings = append(groupStandings, domain.Standing{
                Rank:          t.Rank,
                TeamName:      t.Team.Name,
                TeamLogo:      t.Team.Logo,
                Points:        t.Points,
                GoalsDiff:     t.GoalsDiff,
                MatchesPlayed: t.All.Played,
                Wins:          t.All.Win,
                Draws:         t.All.Draw,
                Losses:        t.All.Lose,
            })
        }
        simplified.Standings = groupStandings
    }

    if err := r.SaveStandings(ctx, leagueID, season, simplified); err != nil {
        fmt.Printf("Warning: could not save to cache: %v\n", err)
    }

    return simplified, nil
}


func (r *StandingsRepo) SaveStandings(ctx context.Context, leagueID, season int, standings *domain.StandingsResponse) error {
	key := fmt.Sprintf("st:%d:%d", leagueID, season)
	payload, err := json.Marshal(standings)
	if err != nil {
		return domain.ErrDuplicateFound
	}

	if err := r.rdb.Set(ctx, key, payload, 0).Err(); err != nil {
		return domain.ErrInternalServer
	}
	return nil
}

func (r *StandingsRepo) GetStandingsFromCache(ctx context.Context, leagueID, season int) (*domain.StandingsResponse, error) {
	key := fmt.Sprintf("st:%d:%d", leagueID, season)
	cachedData, err := r.rdb.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, domain.ErrInternalServer
		}
		return nil, domain.ErrInternalServer
	}

	var standingsResponse domain.StandingsResponse
	if err := json.Unmarshal(cachedData, &standingsResponse); err != nil {
		return nil, domain.ErrInternalServer
	}

	return &standingsResponse, nil
}
