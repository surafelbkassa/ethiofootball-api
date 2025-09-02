package infrastructure

import (
	"fmt"
)

// RedisConnect is just a placeholder right now
// later: configure actual Redis client
func RedisConnect() string {
	return "redis-client-stub"
}

// Example external API call placeholder
// later: wire API-Football or Ethiopian League source
func FetchFixturesFromAPI(league, team, from, to string) []map[string]string {
	fmt.Println("Fetching from API...", league, team, from, to)

	// dummy data
	return []map[string]string{
		{
			"id":      "1",
			"league":  league,
			"home_id": "Chelsea",
			"away_id": "Arsenal",
			"date":    "2025-09-14",
			"status":  "scheduled",
		},
	}
}
