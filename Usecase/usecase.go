package usecase

import (
	"context"
	"errors"
	"log"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
)

type TeamUsecases interface {
	GetTeam(ctx context.Context, teamId string) (*domain.Team, error)
	AddTeam(ctx context.Context, team *domain.Team) error
}

func NewTeamUsecase(repo domain.IRedisRepo) TeamUsecases {
	return &teamUsecase{teamRepo: repo}
}

type teamUsecase struct {
	teamRepo domain.IRedisRepo
}

func (tu *teamUsecase) GetTeam(ctx context.Context, teamId string) (*domain.Team, error) {
	return tu.teamRepo.Get(ctx, teamId)
}

func (tu *teamUsecase) AddTeam(ctx context.Context, team *domain.Team) error {
	return tu.teamRepo.Add(ctx, team)
}

type FixtureUsecase interface {
	GetFixtures(ctx context.Context, league, team, from, to string) ([]domain.Fixture, error)
}

type fixtureUsecase struct {
	repo  repository.FixtureRepo
	cache repository.FixtureRepo
}

func NewFixtureUsecase(r repository.FixtureRepo, c repository.FixtureRepo) FixtureUsecase {
	return &fixtureUsecase{
		repo:  r,
		cache: c,
	}
}

func (uc *fixtureUsecase) GetFixtures(ctx context.Context, league, team, from, to string) ([]domain.Fixture, error) {
	if league == "" {
		return nil, errors.New("league is required")
	}

	// Try cache first
	fixtures, err := uc.cache.GetFixtures(league, team, from, to)
	if err == nil && len(fixtures) > 0 {
		return fixtures, nil
	}

	if err != nil {
		log.Printf("cache miss or error fetching fixtures (league=%s, team=%s, from=%s, to=%s): %v", league, team, from, to, err)
	}

	// Fallback to API repo
	fixtures, err = uc.repo.GetFixtures(league, team, from, to)
	if err != nil {
		log.Printf("API fetch failed (league=%s, team=%s, from=%s, to=%s): %v", league, team, from, to, err)
		return nil, err
	}

	if apiRepo, ok := uc.cache.(*repository.APIRepo); ok && apiRepo.RDB != nil {
		if err := apiRepo.SetFixturesCache(league, team, from, to, fixtures); err != nil {
			log.Printf("failed to cache fixtures: %v", err)
		}
	}
	return fixtures, nil
}
