package domain

// Team represents a football team
type Team struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Short    string `json:"short"`
	League   string `json:"league"`
	CrestURL string `json:"crest_url"`
	Bio      string `json:"bio"`
}

// Standing represents league table information
type Standing struct {
	League      string `json:"league"`
	Season      string `json:"season"`
	Pos         int    `json:"pos"`
	TeamID      string `json:"team_id"`
	Points      int    `json:"pts"`
	GoalDiff    int    `json:"gd"`
	LastUpdated string `json:"last_updated"`
}

// Fixture represents a scheduled or completed match
type Fixture struct {
	ID          string `json:"id"`
	League      string `json:"league"`
	DateUTC     string `json:"date_utc"`
	HomeID      string `json:"home_id"`
	AwayID      string `json:"away_id"`
	Status      string `json:"status"`
	Score       string `json:"score"`
	LastUpdated string `json:"last_updated"`
}

// NewsItem represents a news article
type NewsItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Snippet     string `json:"snippet"`
	Source      string `json:"source"`
	URL         string `json:"url"`
	PublishedAt string `json:"published_at"`
}

// FollowedTeam represents a user's followed team
type FollowedTeam struct {
	TeamID    string `json:"team_id"`
	CreatedAt string `json:"created_at"`
	Notify    bool   `json:"notify"`
}

// CacheMeta represents metadata for cached entries
type CacheMeta struct {
	Key         string `json:"key"`
	Source      string `json:"source"`
	FreshnessTS string `json:"freshness_ts"`
}
