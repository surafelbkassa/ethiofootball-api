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

type APIResponse struct {
	Get        string        `json:"get"`
	Parameters interface{}   `json:"parameters"` 
	Errors     []interface{} `json:"errors"`
	Results    int           `json:"results"`
	Paging     Paging        `json:"paging"`
	Response   []Match       `json:"response"`
}

type Paging struct {
	Current int `json:"current"`
	Total   int `json:"total"`
}

type Match struct {
	Fixture PFixture `json:"fixture"`
	League  League `json:"league"`
	Teams   Teams   `json:"teams"`
	Goals   Goals   `json:"goals"`
	Score   Score `json:"score"`
}

type PFixture struct {
	ID        int     `json:"id"`
	Referee   *string `json:"referee"`
	Timezone  string  `json:"timezone"`
	Date      string  `json:"date"`
	Timestamp int64   `json:"timestamp"`
	Venue     Venue   `json:"venue"`
	Status    Status  `json:"status"`
}

type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed int    `json:"elapsed"`
	Extra   *int   `json:"extra"`
}

type League struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Country   string `json:"country"`
	Logo      string `json:"logo"`
	Flag      string `json:"flag"`
	Season    int    `json:"season"`
	Round     string `json:"round"`
	Standings bool   `json:"standings"`
}

type Teams struct {
	Home MTeam `json:"home"`
	Away MTeam `json:"away"`
}

type MTeam struct {
	Name   string  `json:"name"`
	Logo   string  `json:"logo"`
}

type Goals struct {
	Home *int `json:"home"`
	Away *int `json:"away"`
}

type Score struct {
	Halftime  Goals `json:"halftime"`
	Fulltime  Goals `json:"fulltime"`
	Extratime Goals `json:"extratime"`
	Penalty   Goals `json:"penalty"`
}

type PreviousFixtures struct {
	FixtureID int `json:"fixture_id"`
	Date string `json:"date"`
	Venue string `json:"venue"`
	League string `json:"league"`
	LeagueRound string `json:"round"`
	HomeTeam MTeam `json:"home_team"`
	AwayTeam MTeam `json:"away_team"`
	Goals Goals `json:"goals"`
	Score Score `json:"score"`
	
}
