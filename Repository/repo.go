package repository

import (
	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	infrastructure "github.com/abrshodin/ethio-fb-backend/Infrastructure"
)

// FixtureRepo is an interface to abstract fixture data fetching
type FixtureRepo interface {
	GetFixtures(league, team, from, to string) ([]domain.Fixture, error)
}

// APIRepo implements FixtureRepo using Infrastructure
type APIRepo struct{}

func (r *APIRepo) GetFixtures(league, team, from, to string) ([]domain.Fixture, error) {
	raw := infrastructure.FetchFixturesFromAPI(league, team, from, to)

	var fixtures []domain.Fixture
	for _, item := range raw {
		fixtures = append(fixtures, domain.Fixture{
			ID:      item["id"],
			League:  item["league"],
			HomeID:  item["home_id"],
			AwayID:  item["away_id"],
			DateUTC: item["date"],
			Status:  item["status"],
			Score:   "0-0", // placeholder
		})
	}
	return fixtures, nil
}
