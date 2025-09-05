package usecase

import (
	"fmt"
	"strconv"
	"time"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type EventRepository interface {
	GetPastEvents() ([]domain.Event, error)
	GetStandings() ([]domain.LeaguePoint, error)
	GetFutureEvents() ([]domain.Event, error)
	GetLiveScores() ([]domain.Event, error)
}

type NewsUseCase struct {
	repo EventRepository
}

func NewNewsUseCase(repo EventRepository) *NewsUseCase {
	return &NewsUseCase{repo: repo}
}

// --- Past Events ---
func (uc *NewsUseCase) GenerateNews() ([]string, error) {
	events, err := uc.repo.GetPastEvents()
	if err != nil {
		return nil, err
	}

	var news []string
	for _, e := range events {
		dateStr := e.DateEvent
		if parsedDate, err := time.Parse("2006-01-02", e.DateEvent); err == nil {
			dateStr = parsedDate.Format("02 Jan 2006")
		}

		homeScore, _ := strconv.Atoi(e.IntHomeScore)
		awayScore, _ := strconv.Atoi(e.IntAwayScore)

		var resultDesc string
		switch {
		case homeScore == awayScore:
			resultDesc = fmt.Sprintf("%s and %s played to a %d-%d draw", e.StrHomeTeam, e.StrAwayTeam, homeScore, awayScore)
		case homeScore > awayScore:
			resultDesc = fmt.Sprintf("%s edged out %s with a %d-%d victory", e.StrHomeTeam, e.StrAwayTeam, homeScore, awayScore)
		default:
			resultDesc = fmt.Sprintf("%s dominated %s with a %d-%d win", e.StrAwayTeam, e.StrHomeTeam, awayScore, homeScore)
		}

		headline := fmt.Sprintf("%s on %s | Status: %s", resultDesc, dateStr, e.StrStatus)
		news = append(news, headline)
	}

	return news, nil
}

// --- Standings ---
func (uc *NewsUseCase) GenerateStandingNews() ([]string, error) {
	standings, err := uc.repo.GetStandings()
	if err != nil {
		return nil, err
	}

	var news []string
	for _, s := range standings {
		headline := fmt.Sprintf(
			"%s is ranked #%s with %s points (%s Wins - %s Draws - %s Losses).",
			s.StrTeam,
			s.IntRank,
			s.IntPoints,
			s.IntWin,
			s.IntDraw,
			s.IntLoss,
		)
		news = append(news, headline)
	}

	return news, nil
}

// --- Future Events ---
func (uc *NewsUseCase) GenerateFutureNews() ([]string, error) {
	events, err := uc.repo.GetFutureEvents()
	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return []string{"No upcoming events or games at the moment. Stay tuned!"}, nil
	}

	var news []string
	for _, e := range events {
		dateStr := e.DateEvent
		if parsedDate, err := time.Parse("2006-01-02", e.DateEvent); err == nil {
			dateStr = parsedDate.Format("02 Jan 2006")
		}

		headline := fmt.Sprintf(
			"Upcoming showdown: %s vs %s on %s | Status: %s",
			e.StrHomeTeam, e.StrAwayTeam, dateStr, e.StrStatus,
		)
		news = append(news, headline)
	}

	return news, nil
}

// --- Live Scores ---
func (uc *NewsUseCase) GenerateLiveScores() ([]string, error) {
	events, err := uc.repo.GetLiveScores()
	if err != nil {
		return []string{"No live games at the moment. Check back later!"}, nil
	}

	var news []string
	today := time.Now().Format("02 Jan 2006")

	for _, e := range events {
		headline := fmt.Sprintf(
			"Currently playing Live: %s vs %s on %s with score %s to %s | Status: %s, Who will win?",
			e.StrHomeTeam, e.StrAwayTeam, today, e.IntHomeScore, e.IntAwayScore, e.StrStatus,
		)
		news = append(news, headline)
	}

	return news, nil
}

// // --- Optional Aggregator ---
// func (uc *NewsUseCase) GenerateAllNews() ([]string, error) {
// 	var allNews []string

// 	if past, err := uc.GenerateNews(); err == nil {
// 		allNews = append(allNews, past...)
// 	}
// 	if live, err := uc.GenerateLiveScores(); err == nil {
// 		allNews = append(allNews, live...)
// 	}
// 	if future, err := uc.GenerateFutureNews(); err == nil {
// 		allNews = append(allNews, future...)
// 	}

// 	return allNews, nil
// }
