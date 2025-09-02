package usecase

import (
	"fmt"
	"strconv"
	"time"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
)

type EventRepository interface {
	GetPastEvents() ([]domain.Event, error)
}

type NewsUseCase struct {
	repo EventRepository
}

func NewNewsUseCase(repo EventRepository) *NewsUseCase {
	return &NewsUseCase{repo: repo}
}

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

		// Convert string scores to int
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
