package usecase

import (
	"context"
	"errors"
	"strings"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
)

type IFixturesUsecase interface {
	FetchAndStore(ctx context.Context, league string, leagueID int, q domain.RoundQuery) (*[]domain.PrevFixtures, error)
	GetCachedByRound(ctx context.Context, q domain.RoundQuery) (*[]domain.PrevFixtures, error)
	ResolveRoundWindow(ctx context.Context, q domain.RoundQuery) (domain.RoundQuery, error)
	 GetLiveMatches (league string)(*[]domain.PrevFixtures, error)
}

func NewFixturesUsecase(api domain.IAPIService, repo repository.IFixturesRepo) IFixturesUsecase {
	return &FixturesUsecase{api: api, repo: repo}
}

type FixturesUsecase struct {
	api  domain.IAPIService
	repo repository.IFixturesRepo
}

func (uc *FixturesUsecase) FetchAndStore(ctx context.Context, league string, leagueID int, q domain.RoundQuery) (*[]domain.PrevFixtures, error) {

	fixtures, err := uc.api.PrevFixtures(leagueID, q.Season, q.From, q.To)
	if err != nil {
		return nil, err
	}

	// group by round and store
	rounds := map[string][]domain.PrevFixtures{}
	for _, f := range *fixtures {
		round := normalizeRound(f.LeagueRound)
		rounds[round] = append(rounds[round], f)
	}

	for round, fs := range rounds {
		rq := q
		rq.Round = round
		if err := uc.repo.SaveFixturesByRound(ctx, rq, fs); err != nil {
			return nil, err
		}
	}

	// stores round with from, to dates
	for round := range rounds {
		rq := q
		rq.Round = round
		if err := uc.repo.SaveRoundWindow(ctx, rq); err != nil {
			return nil, err
		}
	}

	return fixtures, nil
}

func (uc *FixturesUsecase) GetCachedByRound(ctx context.Context, q domain.RoundQuery) (*[]domain.PrevFixtures, error) {
	fixtures, err := uc.repo.GetFixturesByRound(ctx, q)
	if err != nil {
		return nil, err
	}
	if fixtures == nil || len(*fixtures) == 0 {
		return nil, errors.New("not found")
	}
	return fixtures, nil
}

func normalizeRound(s string) string {

	s = strings.TrimSpace(s)
	parts := strings.Fields(s)
	if len(parts) == 0 {
		return s
	}

	last := parts[len(parts)-1]
	return last
}

func (uc *FixturesUsecase) ResolveRoundWindow(ctx context.Context, q domain.RoundQuery) (domain.RoundQuery, error) {
	if q.From != "" && q.To != "" {
		return q, nil
	}

	if from, to, err := uc.repo.GetRoundWindow(ctx, q); err == nil && from != "" && to != "" {
		q.From, q.To = from, to
		return q, nil
	}

	seasonWindows := map[string]map[int]struct{ From, To string }{
		"ETH": {
			2021: {From: "2021-10-17", To: "2022-06-28"},
			2022: {From: "2022-09-30", To: "2023-07-08"},
			2023: {From: "2023-10-01", To: "2024-06-30"},
		},

		"EPL": {
			2021: {From: "2021-08-24", To: "2022-04-03"},
			2022: {From: "2022-08-30", To: "2023-04-02"},
			2023: {From: "2023-08-22", To: "2024-04-07"},
		},
	}

	if seasons, ok := seasonWindows[q.League]; ok {
		if win, ok := seasons[q.Season]; ok {
			q.From, q.To = win.From, win.To
			_ = uc.repo.SaveRoundWindow(ctx, q)
			return q, nil
		}
	}

	return q, errors.New("round window not found")
}

func(uc *FixturesUsecase) GetLiveMatches (league string)(*[]domain.PrevFixtures, error){
	return uc.api.LiveFixtures(league)
}
