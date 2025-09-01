package repository

import (
	"context"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	"github.com/redis/go-redis/v9"
)

func NewTeamRepo (rdb *redis.Client) domain.TeamRepo {
	return &teamRepo{rdb: rdb}
}
type teamRepo struct {
	rdb *redis.Client
}

func(tr *teamRepo) Get(ctx context.Context, teamId string) (*domain.Team, error){

	key := "team" + teamId
	vals, err := tr.rdb.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, domain.ErrInternalServer
	}

	if vals == nil {
		return nil, domain.ErrTeamNotFound
	}

	team := &domain.Team{
		ID: vals["id"],
		Name: vals["name"],
		Short: vals["short"],
		League: vals["league"],
		CrestURL: vals["crest_url"],
		Bio: vals["bio"],
	}
	return team, nil
}

func(tr *teamRepo) Add(ctx context.Context, team *domain.Team) error{

	key := "team" + team.ID
	exists, err := tr.rdb.Exists(ctx, key).Result()

	if exists > 0 {
		return domain.ErrDuplicateFound
	}
	err = tr.rdb.HSet(ctx, key, map[string]interface{}{
		"ID": team.ID,
		"name": team.Name,
		"short": team.Short,
		"league": team.League,
		"crest_url": team.CrestURL,
		"bio": team.Bio,
		}).Err()

	if err != nil {
		return domain.ErrInternalServer
	}

	return nil
}

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
