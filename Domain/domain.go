package domain

// Team represents a football team
type Team struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Short    string    `json:"short"`
	League   string    `json:"league"`
	CrestURL string    `json:"crest_url"`
	Bio      string    `json:"bio"`
	Fixtures []Fixture     
}

// Fixture represents a scheduled or completed match

type Fixture struct {
	ID          string `json:"id"`
	League      string `json:"league"`
	DateUTC     string `json:"date_utc"` // ISO string
	HomeID      string `json:"home_id"`
	AwayID      string `json:"away_id"`
	Status      string `json:"status"`
	Score       string `json:"score"`
	HomeLogo    string `json:"home_logo,omitempty"`
	AwayLogo    string `json:"away_logo,omitempty"`
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
	Response []Match `json:"response"`
}

type Match struct {
	Fixture PFixture `json:"fixture"`
	League  League   `json:"league"`
	Teams   Teams    `json:"teams"`
	Goals   Goals    `json:"goals"`
	Score   Score    `json:"score"`
	
}

type PFixture struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Timestamp int64  `json:"timestamp"`
	Venue     Venue  `json:"venue"`
	Status    Status `json:"status"`
	Played 	  TeamStats `json:"played"`
	Wins 	  TeamStats `json:"wins"`
	Lose	  TeamStats	`json:"loses"`
	Draws 	  TeamStats	`json:"draws"`

}

type Venue struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type League struct {
	ID   		int    				`json:"id"`
	Name    	string 				`json:"name"`
	Country 	string 				`json:"country"`
	Round   	string 				`json:"round"`
	Logo     	string      		`json:"logo"`
	Flag     	string      		`json:"flag"`
	Season   	int         		`json:"season"`
	Standings 	[][] TeamStanding   `json:"standings"`
}

type Teams struct {
	Home MTeam `json:"home"`
	Away MTeam `json:"away"`
}

type MTeam struct {
	Name string `json:"name"`
	Logo string `json:"logo"`
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

type EventTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

type EventPlayer struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type EventAssist struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type Status struct {
	Long    string `json:"long"`
	Short   string `json:"short"`
	Elapsed int    `json:"elapsed"`
	Extra   int    `json:"extra"`
}

type PrevFixtures struct {
	FixtureID   int    `json:"fixture_id"`
	Date        string `json:"date"`
	Venue       string `json:"venue"`
	League      string `json:"league"`
	LeagueRound string `json:"round"`
	HomeTeam    MTeam  `json:"home_team"`
	AwayTeam    MTeam  `json:"away_team"`
	Goals       Goals  `json:"goals"`
	Score       Score  `json:"score"`
	Status      Status `json:"status"`
}

type RoundQuery struct {
	League string `json:"league"`
	Season int    `json:"season"`
	Round  string `json:"round"`
	From   string `json:"from"`
	To     string `json:"to"`
}

// Event represents a football match event
type Event struct {
	IDEvent           string `json:"idEvent"`
	StrEvent          string `json:"strEvent"`
	StrEventAlternate string `json:"strEventAlternate"`
	StrLeague         string `json:"strLeague"`
	StrSeason         string `json:"strSeason"`
	StrHomeTeam       string `json:"strHomeTeam"`
	StrAwayTeam       string `json:"strAwayTeam"`
	IntHomeScore      string `json:"intHomeScore"`
	IntAwayScore      string `json:"intAwayScore"`
	DateEvent         string `json:"dateEvent"`
	StrTime           string `json:"strTime"`
	StrStatus         string `json:"strStatus"`
	StrHomeTeamBadge  string `json:"strHomeTeamBadge"`
	StrAwayTeamBadge  string `json:"strAwayTeamBadge"`
}

type LeaguePoint struct {
	IDStanding        string `json:"idStanding"`
	IntRank           string `json:"intRank"`
	IDTeam            string `json:"idTeam"`
	StrTeam           string `json:"strTeam"`
	StrBadge          string `json:"strBadge"`
	IDLeague          string `json:"idLeague"`
	StrLeague         string `json:"strLeague"`
	StrSeason         string `json:"strSeason"`
	StrForm           string `json:"strForm"`
	StrDescription    string `json:"strDescription"`
	IntPlayed         string `json:"intPlayed"`
	IntWin            string `json:"intWin"`
	IntLoss           string `json:"intLoss"`
	IntDraw           string `json:"intDraw"`
	IntGoalsFor       string `json:"intGoalsFor"`
	IntGoalsAgainst   string `json:"intGoalsAgainst"`
	IntGoalDifference string `json:"intGoalDifference"`
	IntPoints         string `json:"intPoints"`
	DateUpdated       string `json:"dateUpdated"`
}
type StandingAPIResponse struct {
	Response []Response `json:"response"`
}

type Response struct{
	League League `json:"league"`
}

type TeamStanding struct {
    Rank          int          `json:"rank"`
    Team          StandingTeam `json:"team"`
    Points        int          `json:"points"`
    GoalsDiff     int          `json:"goalsDiff"`
    MatchesPlayed int          `json:"matchesPlayed"`
    Wins          int          `json:"wins"`
    Losses        int          `json:"losses"`
    Draws         int          `json:"draws"`
    All           MatchStats   `json:"all"`
}

type StandingTeam struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Logo string `json:"logo"`
}

type MatchStats struct {
    Played int       `json:"played"`
    Win    int       `json:"win"`
    Draw   int       `json:"draw"`
    Lose   int       `json:"lose"`
    Goals  GoalStats `json:"goals"`
}

type GoalStats struct {
    For     int `json:"for"`
    Against int `json:"against"`
}

type Standing struct {
    Rank          int    `json:"rank"`
    TeamName      string `json:"teamName"`
    TeamLogo      string `json:"teamLogo"`
    Points        int    `json:"points"`
    MatchesPlayed int    `json:"matchesPlayed"`
    Wins          int    `json:"wins"`
    Losses        int    `json:"losses"`
    Draws         int    `json:"draws"`
    GoalsDiff     int    `json:"goalsDiff"`
}

type StandingsResponse struct {
    LeagueID    int        `json:"leagueId"`
    LeagueName  string     `json:"leagueName"`
    Country     string     `json:"country"`
    CountryFlag string     `json:"countryFlag"`
    Season      int        `json:"season"`
    Standings   []Standing `json:"standings"`
    LastUpdated string     `json:"lastUpdated"`
}

type StatAPIResponse struct {
	Response  Stats `json:"response"`
}

type Stats struct {
	League League 	`json:"league"`
	Team   MTeam	`json:"team"`
	Fixture PFixture `json:"fixture"`
	Goals   GoalStat    `json:"goals"`
}

type TeamStats struct {
	Home 	int `json:"home"`
	Away 	int `json:"away"`
	Total 	int `json:"total"`
}

type GoalStat struct {
	For Total `json:"for"`
	Against Total `json:"against"`
}

type Total struct{
	Total TeamStats `json:"total"`
}




