package usecase

import (
	"context"
	"fmt"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type IStandingsUsecase interface {
	GetStandings(ctx context.Context, leagueID, season int) (*domain.StandingsResponse, error)
}

type StandingsUsecase struct {
	standingsRepo domain.IStandingsRepo
}

func NewStandingsUsecase(standingsRepo domain.IStandingsRepo) IStandingsUsecase {
	return &StandingsUsecase{
		standingsRepo: standingsRepo,
	}
}

func (u *StandingsUsecase) GetStandings(ctx context.Context, leagueID, season int) (*domain.StandingsResponse, error) {

	standings, err := u.standingsRepo.GetStandings(ctx, leagueID, season)
	if err != nil {
		return nil, fmt.Errorf("failed to get standings: %w", err)
	}

	return standings, nil
}
