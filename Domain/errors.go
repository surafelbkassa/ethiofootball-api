package domain

import "errors"

var (
	ErrInternalServer = errors.New("internal server error")
	ErrDuplicateFound = errors.New("duplicate key found")
	ErrTeamNotFound   = errors.New("team not found")
	ErrUnexpected     = errors.New("Unexpected")
)
