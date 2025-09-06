package usecase

import (
	"github.com/abrshodin/ethio-fb-backend/Domain"
)

type IntentParser interface {
	Parse(txt string) (*domain.Intent, error)
}

type ParseIntentUseCase struct {
	parser IntentParser
}

// NewParseIntentUsecase creates a new ParseIntentUseCase with the given IntentParser.
func NewParseIntentUsecase(parser IntentParser) *ParseIntentUseCase {
	return &ParseIntentUseCase{parser}
}

// Execute transforms the given text into an Intent object using the configured Parser.
func (uc *ParseIntentUseCase) Execute(text string) (*domain.Intent, error) {
	if text == "" {
		return nil, ErrInvalidInput
	}
	return uc.parser.Parse(text)
}
