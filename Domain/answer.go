package domain

import (
	"time"
)

type Answer struct {
	Markdown  string    `json:"markdown"`
	Source    string    `json:"source"`
	Freshness time.Time `json:"freshness"`
}

type AnswerContext struct {
	Language    string
	Source      string
	Freshness   time.Time
	ContextData map[string]interface{}
}

type AnswerComposer interface {
	ComposeAnswer(context AnswerContext) (*Answer, error)
}
