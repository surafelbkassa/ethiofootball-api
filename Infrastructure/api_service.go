package infrastructure

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type IAPIService interface {
	PrevFixtures(leagueID int, season int, fromDate, toDate string) (*[]domain.PrevFixtures, error)
	LiveFixtures(league string) (*[]domain.PrevFixtures, error)
}

func NewAPIService() IAPIService {
	return &APIServiceClient{}
}

type APIServiceClient struct{}

func (hs APIServiceClient) PrevFixtures(leagueID int, season int, fromDate, toDate string) (*[]domain.PrevFixtures, error) {

	API_KEY := os.Getenv("API_SPORTS_API_KEY")

	url := fmt.Sprintf(
		"https://v3.football.api-sports.io/fixtures?league=%d&season=%d&from=%s&to=%s",
		leagueID, season, fromDate, toDate,
	)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, domain.ErrInternalServer
	}

	// API headers
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

	var apiResponse domain.APIResponse
	if err := json.Unmarshal([]byte(string(body)), &apiResponse); err != nil {
		return nil, domain.ErrInternalServer
	}

	var prevFixtures = &[]domain.PrevFixtures{}
	for _, r := range apiResponse.Response {
		fixture := domain.PrevFixtures{
			Date:        r.Fixture.Date,
			Venue:       r.Fixture.Venue.Name,
			League:      r.League.Name,
			LeagueRound: r.League.Round,
			HomeTeam: domain.MTeam{
				Name: r.Teams.Home.Name,
				Logo: r.Teams.Home.Logo,
			},

			AwayTeam: domain.MTeam{
				Name: r.Teams.Away.Name,
				Logo: r.Teams.Away.Logo,
			},

			Goals: domain.Goals{
				Home: r.Goals.Home,
				Away: r.Goals.Away,
			},

			Score: domain.Score{
				Halftime:  domain.Goals(r.Score.Halftime),
				Fulltime:  domain.Goals(r.Score.Fulltime),
				Extratime: domain.Goals(r.Score.Extratime),
				Penalty:   domain.Goals(r.Score.Penalty),
			},
		}

		*prevFixtures = append(*prevFixtures, fixture)
	}

	return prevFixtures, nil
}

func (ac *APIServiceClient) LiveFixtures(league string) (*[]domain.PrevFixtures, error) {

	ids := ""
	if league == "ETH" {
		ids = "363-363"
	} else {
		ids = "39-39"
	}

	API_KEY := os.Getenv("API_SPORTS_API_KEY")

	url := fmt.Sprintf("https://v3.football.api-sports.io/fixtures?live=%s", ids)

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, domain.ErrInternalServer
	}

	// API headers
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

	var apiResponse domain.APIResponse
	if err := json.Unmarshal([]byte(string(body)), &apiResponse); err != nil {
		return nil, domain.ErrInternalServer
	}

	if len(apiResponse.Response) == 0 {
		return nil, nil
	}

	var prevFixtures = &[]domain.PrevFixtures{}

	
	for _, r := range apiResponse.Response {

		fixture := domain.PrevFixtures{
			Date:        r.Fixture.Date,
			Venue:       r.Fixture.Venue.Name,
			League:      r.League.Name,
			LeagueRound: r.League.Round,
			HomeTeam: domain.MTeam{
				Name: r.Teams.Home.Name,
				Logo: r.Teams.Home.Logo,
			},

			AwayTeam: domain.MTeam{
				Name: r.Teams.Away.Name,
				Logo: r.Teams.Away.Logo,
			},

			Goals: domain.Goals{
				Home: r.Goals.Home,
				Away: r.Goals.Away,
			},

			Score: domain.Score{
				Halftime:  domain.Goals(r.Score.Halftime),
				Fulltime:  domain.Goals(r.Score.Fulltime),
				Extratime: domain.Goals(r.Score.Extratime),
				Penalty:   domain.Goals(r.Score.Penalty),
			},
			
			Status: r.Fixture.Status,
		}

		*prevFixtures = append(*prevFixtures, fixture)
	}

	return prevFixtures, nil

}




