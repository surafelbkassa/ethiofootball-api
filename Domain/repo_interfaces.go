package domain

import (
	"context"
)

type IRedisRepo interface {
	Get(ctx context.Context, teamId string) (*Team, error)
	Add(ctx context.Context, team *Team) error
}

type IStandingsRepo interface {
	GetStandings(ctx context.Context, leagueID, season int) (*StandingsResponse, error)
	SaveStandings(ctx context.Context, leagueID, season int, standings *StandingsResponse) error
	GetStandingsFromCache(ctx context.Context, leagueID, season int) (*StandingsResponse, error)
}
