package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type EventRepositoryImpl struct {
	apiURL         string
	apiStandingURL string
	apiFutureURL   string
}

func NewEventRepository() *EventRepositoryImpl {
	return &EventRepositoryImpl{
		apiURL:         "https://www.thesportsdb.com/api/v1/json/123/eventspastleague.php?id=4959",
		apiStandingURL: "https://www.thesportsdb.com/api/v1/json/123/lookuptable.php?l=4959",
		apiFutureURL:   "https://www.thesportsdb.com/api/v1/json/123/eventsnextleague.php?id=4959",
	}
}

func (r *EventRepositoryImpl) GetPastEvents() ([]domain.Event, error) {
	resp, err := http.Get(r.apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result struct {
		Events []domain.Event `json:"events"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return result.Events, nil
}

func (r *EventRepositoryImpl) GetStandings() ([]domain.LeaguePoint, error) {
	resp, err := http.Get(r.apiStandingURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch standings: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result struct {
		Table []domain.LeaguePoint `json:"table"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal standings JSON: %v", err)
	}

	return result.Table, nil
}

func (r *EventRepositoryImpl) GetFutureEvents() ([]domain.Event, error) {
	resp, err := http.Get(r.apiFutureURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch events: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result struct {
		Events []domain.Event `json:"events"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	// Handle null safely (If future games are not decided)
	if result.Events == nil {
		return []domain.Event{}, nil
	}

	return result.Events, nil
}

// --- Demo Live Scores until API is found ---
var demoLiveScores = []domain.Event{
	{
		IDEvent:           "1001",
		StrEvent:          "Saint George vs Ethiopian Coffee",
		StrEventAlternate: "St George v Coffee",
		StrLeague:         "Ethiopian Premier League",
		StrSeason:         "2024/2025",
		StrHomeTeam:       "Saint George",
		StrAwayTeam:       "Ethiopian Coffee",
		IntHomeScore:      "2",
		IntAwayScore:      "1",
		DateEvent:         "2025-09-05",
		StrTime:           "15:30",
		StrStatus:         "Live - 67'",
		StrHomeTeamBadge:  "https://example.com/badges/saint_george.png",
		StrAwayTeamBadge:  "https://example.com/badges/ethiopian_coffee.png",
	},
	{
		IDEvent:           "1002",
		StrEvent:          "Adama City vs Fasil Kenema",
		StrEventAlternate: "Adama v Fasil",
		StrLeague:         "Ethiopian Premier League",
		StrSeason:         "2024/2025",
		StrHomeTeam:       "Adama City",
		StrAwayTeam:       "Fasil Kenema",
		IntHomeScore:      "0",
		IntAwayScore:      "0",
		DateEvent:         "2025-09-05",
		StrTime:           "16:00",
		StrStatus:         "Live - HT",
		StrHomeTeamBadge:  "https://example.com/badges/adama.png",
		StrAwayTeamBadge:  "https://example.com/badges/fasil.png",
	},
	{
		IDEvent:           "1003",
		StrEvent:          "Sidama Bunna vs Hawassa City",
		StrEventAlternate: "Sidama v Hawassa",
		StrLeague:         "Ethiopian Premier League",
		StrSeason:         "2024/2025",
		StrHomeTeam:       "Sidama Bunna",
		StrAwayTeam:       "Hawassa City",
		IntHomeScore:      "1",
		IntAwayScore:      "3",
		DateEvent:         "2025-09-05",
		StrTime:           "14:00",
		StrStatus:         "Live - 80'",
		StrHomeTeamBadge:  "https://example.com/badges/sidama.png",
		StrAwayTeamBadge:  "https://example.com/badges/hawassa.png",
	},
}

func (r *EventRepositoryImpl) GetLiveScores() ([]domain.Event, error) {
	if len(demoLiveScores) == 0 {
		return []domain.Event{}, fmt.Errorf("no live scores available right now")
	}
	return demoLiveScores, nil
}
