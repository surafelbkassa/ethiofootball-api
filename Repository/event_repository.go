package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type EventRepositoryImpl struct {
	apiURL string
}

func NewEventRepository() *EventRepositoryImpl {
	return &EventRepositoryImpl{
		apiURL: "https://www.thesportsdb.com/api/v1/json/123/eventspastleague.php?id=4959",
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
