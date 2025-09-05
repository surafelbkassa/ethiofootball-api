package usecase

import (
  "context"
	"errors"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
)
// team usecase
type TeamUsecases interface {
	GetTeam(ctx context.Context, teamId string) (*domain.Team, error)
	AddTeam(ctx context.Context, team *domain.Team) error
}

func NewTeamUsecase(repo domain.IRedisRepo) TeamUsecases{
	return &teamUsecase{teamRepo: repo}
}

type teamUsecase struct{
	teamRepo domain.IRedisRepo
}

func(tu *teamUsecase) GetTeam(ctx context.Context, teamId string)(*domain.Team, error){
	return tu.teamRepo.Get(ctx, teamId)
} 

func(tu *teamUsecase) AddTeam(ctx context.Context, team *domain.Team) error{
	return tu.teamRepo.Add(ctx, team)
}

// fixture
type FixtureUsecase interface {
	GetFixtures(league, team, from, to string) ([]domain.Fixture, error)
}

type fixtureUsecase struct {
	repo repository.FixtureRepo
}

func NewFixtureUsecase(r repository.FixtureRepo) FixtureUsecase {
	return &fixtureUsecase{repo: r}
}

func (uc *fixtureUsecase) GetFixtures(league, team, from, to string) ([]domain.Fixture, error) {
	if league == "" {
		return nil, errors.New("league is required")
	}
	// TODO: validate date format, team existence

	return uc.repo.GetFixtures(league, team, from, to)
}

package usecase

import (
	"context"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type AnswerUsecase interface {
	Compose(ctx context.Context, context domain.AnswerContext) (*domain.Answer, error)
}
