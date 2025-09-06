package domain

import "time"

type ComparisonTeam struct {
	Name           string   `json:"name"`
	Honors         []string `json:"honors"`
	RecentForm     []string `json:"recent_form"`
	NotablePlayers []string `json:"notable_players"`
	FanbaseNotes   string   `json:"fanbase_notes"`
}

type ComparisonData struct {
	TeamA ComparisonTeam `json:"team_a"`
	TeamB ComparisonTeam `json:"team_b"`
}

type Answer struct {
	Markdown       string          `json:"markdown,omitempty"`
	ComparisonData *ComparisonData `json:"comparison_data,omitempty"`
	Source         string          `json:"source"`
	Freshness      time.Time       `json:"freshness"`
}

type AnswerContext struct {
	Topic       string // "compare", "fixture", "table", etc.
	Language    string
	Source      string
	Freshness   time.Time
	ContextData map[string]interface{}
}

type AnswerComposer interface {
	ComposeAnswer(context AnswerContext) (*Answer, error)
}
