package usecase

import (
	"context"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type TeamUsecases interface {
	GetTeam(ctx context.Context, teamId string) (*domain.Team, error)
	AddTeam(ctx context.Context, team *domain.Team) error
}

func NewTeamUsecase(repo domain.TeamRepo) TeamUsecases{
	return &teamUsecase{teamRepo: repo}
}

type teamUsecase struct{
	teamRepo domain.TeamRepo
}

func(tu *teamUsecase) GetTeam(ctx context.Context, teamId string)(*domain.Team, error){
	return tu.teamRepo.Get(ctx, teamId)
} 

func(tu *teamUsecase) AddTeam(ctx context.Context, team *domain.Team) error{
	return tu.teamRepo.Add(ctx, team)
}




