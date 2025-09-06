package usecase

import (
	"context"
	"errors"

	// "log"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
)

type TeamUsecases interface {
	GetTeam(ctx context.Context, teamId string) (*domain.Team, error)
	AddTeam(ctx context.Context, team *domain.Team) error
	Statistics(ctx context.Context, league, season int, team string) (*domain.TeamComparison, error)
}

func NewTeamUsecase(repo domain.IRedisRepo, api domain.IAPIService) TeamUsecases {
	return &TeamUsecase{teamRepo: repo, api: api}
}

type TeamUsecase struct {
	teamRepo domain.IRedisRepo
	api domain.IAPIService
}

func (tu *TeamUsecase) GetTeam(ctx context.Context, teamId string) (*domain.Team, error) {
	return tu.teamRepo.Get(ctx, teamId)
}

func (tu *TeamUsecase) AddTeam(ctx context.Context, team *domain.Team) error {
	return tu.teamRepo.Add(ctx, team)
}

func (tu *TeamUsecase) Statistics(ctx context.Context, league, season int, team string) (*domain.TeamComparison, error){
	teamID := 1001
	if team == "EthiopianCoffee" {
		teamID = 1004
	} 
	return tu.api.Statistics(league, season, teamID)
}

type FixtureUsecase interface {
	GetFixtures(ctx context.Context, league, team, season, from, to string) ([]domain.Fixture, error)
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

// func (uc *fixtureUsecase) GetFixtures(ctx context.Context, league, team, season, from, to string) ([]domain.Fixture, error) {
// 	if league == "" {
// 		return nil, errors.New("league is required")
// 	}

// 	// Try cache first
// 	fixtures, err := uc.cache.GetFixtures(league, team, season, from, to)
// 	if err == nil && len(fixtures) > 0 {
// 		return fixtures, nil
// 	}

// 	if err != nil {
// 		log.Printf("cache miss or error fetching fixtures (league=%s, team=%s, from=%s, to=%s): %v", league, team, from, to, err)
// 	}

// 	// Fallback to API repo
// 	fixtures, err = uc.repo.GetFixtures(league, team, season, from, to)
// 	if err != nil {
// 		log.Printf("API fetch failed (league=%s, team=%s, from=%s, to=%s): %v", league, team, from, to, err)
// 		return nil, err
// 	}

// 	if apiRepo, ok := uc.cache.(*repository.APIRepo); ok && apiRepo.RDB != nil {
// 		if err := apiRepo.SetFixturesCache(league, team, season, from, to, fixtures); err != nil {
// 			log.Printf("failed to cache fixtures: %v", err)
// 		}
// 	}
// 	return fixtures, nil
// }

func (uc *fixtureUsecase) GetFixtures(ctx context.Context, league, team, season, from, to string) ([]domain.Fixture, error) {
	if league == "" {
		return nil, errors.New("league is required")
	}

	// Call repo once and use its result
	fixtures, err := uc.repo.GetFixtures(league, team, season, from, to)
	if err != nil {
		// propagate error (repo may return non-nil err on auth/network issues)
		return nil, err
	}
	// always return empty slice instead of nil to avoid JSON "null"
	if fixtures == nil {
		return []domain.Fixture{}, nil
	}
	return fixtures, nil
}


type AnswerUsecase interface {
	Compose(ctx context.Context, context domain.AnswerContext) (*domain.Answer, error)
}
