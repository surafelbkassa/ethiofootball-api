package domain

import "context"

type IRedisRepo interface {
	Get(ctx context.Context, teamId string) (*Team, error)
	Add(ctx context.Context,  team *Team) error
}