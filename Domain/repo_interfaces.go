package domain

import "context"

type TeamRepo interface {
	Get(ctx context.Context, teamId string) (*Team, error)
	Add(ctx context.Context,  team *Team) error
}