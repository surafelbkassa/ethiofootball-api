package usecase

import "errors"

// Semantic errors for the interface layer
var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrIntentNotFound     = errors.New("could not parse intent")
	ErrServiceUnavailable = errors.New("intent service unavailable")
)
