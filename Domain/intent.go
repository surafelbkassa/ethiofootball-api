package domain

type Intent struct {
	Topic    string   `json:"topic"`
	Teams    []string `json:"teams"`
	League   string   `json:"league"`
	Date     string   `json:"date,omitempty"`
	FollowUp string   `json:"follow_up,omitempty"`
	Language string   `json:"language"`
}
