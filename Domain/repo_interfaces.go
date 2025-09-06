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

type IAPIService interface {
	PrevFixtures(leagueID int, season int, fromDate, toDate string) (*[]PrevFixtures, error)
	LiveFixtures(league string) (*[]PrevFixtures, error)
	Statistics(league, season, team int) (*TeamComparison, error)
}
