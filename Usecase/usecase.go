package usecase

import (
	"errors"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
)

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
