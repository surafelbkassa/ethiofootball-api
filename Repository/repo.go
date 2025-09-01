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
