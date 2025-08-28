package usecase

import (
	"context"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type AnswerUsecase interface {
	Compose(ctx context.Context, context domain.AnswerContext) (*domain.Answer, error)
}
