package usecase

import (
	"context"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type answerUseCase struct {
	composer domain.AnswerComposer
}

func NewAnswerUseCase(composer domain.AnswerComposer) AnswerUsecase {
	return &answerUseCase{
		composer: composer,
	}
}

func (uc *answerUseCase) Compose(ctx context.Context, context domain.AnswerContext) (*domain.Answer, error) {
	if len(context.ContextData) == 0 {
		return nil, ErrInvalidInput
	}
	answer, err := uc.composer.ComposeAnswer(context)
	if err != nil {
		return nil, err
	}

	return answer, nil
}
