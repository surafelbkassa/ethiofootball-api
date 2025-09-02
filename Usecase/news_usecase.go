// usecases/news_usecase.go
package usecase

import (
	"fmt"
	"time"

	repository "github.com/abrshodin/ethio-fb-backend/Repository"
)

type NewsUseCase struct {
	repo *repository.EventRepository
}

func NewNewsUseCase(repo *repository.EventRepository) *NewsUseCase {
	return &NewsUseCase{repo: repo}
}

func (uc *NewsUseCase) GenerateNews() ([]string, error) {
	events, err := uc.repo.GetPastEvents()
	if err != nil {
		return nil, err
	}

	var news []string
	for _, e := range events {
		// Parse and format date nicely
		dateStr := e.DateEvent
		if parsedDate, err := time.Parse("2006-01-02", e.DateEvent); err == nil {
			dateStr = parsedDate.Format("02 Jan 2006")
		}

		// Create a more descriptive headline
		var resultDesc string
		if e.IntHomeScore == e.IntAwayScore {
			resultDesc = fmt.Sprintf("%s and %s played to a %d-%d draw", e.StrHomeTeam, e.StrAwayTeam, e.IntHomeScore, e.IntAwayScore)
		} else if e.IntHomeScore > e.IntAwayScore {
			resultDesc = fmt.Sprintf("%s edged out %s with a %d-%d victory", e.StrHomeTeam, e.StrAwayTeam, e.IntHomeScore, e.IntAwayScore)
		} else {
			resultDesc = fmt.Sprintf("%s dominated %s with a %d-%d win", e.StrAwayTeam, e.StrHomeTeam, e.IntAwayScore, e.IntHomeScore)
		}

		headline := fmt.Sprintf(
			"%s on %s | Status: %s",
			resultDesc, dateStr, e.StrStatus,
		)
		news = append(news, headline)
	}

	return news, nil
}
