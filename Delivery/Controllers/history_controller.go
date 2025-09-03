package controlller

import (
	"fmt"
	"net/http"
	"strconv"

	infrastructure "github.com/abrshodin/ethio-fb-backend/Infrastructure"
	"github.com/gin-gonic/gin"
)


type HistoryController struct {
	apiService infrastructure.IHistoryAPIService
}

func NewHistoryController(as infrastructure.IHistoryAPIService) *HistoryController {
	return &HistoryController{apiService: as}
}

func(hc *HistoryController) PreviousMatchHistory(c *gin.Context){

	league := c.Query("league")
	from := c.Query("from")
	to := c.Query("to")
	seasonQuery := c.Query("season")

	fmt.Println("league", "from", "to", league, from, to )

	if league == "" || from == "" || to == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "insufficent queries"})
		c.Abort()
		return
	}

	if league != "EPL" && league != "ETH" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid league queries"})
		c.Abort()
		return
	}

	getSeason := func(defaultSeason int) int {
			if seasonQuery != "" {
				if season, err := strconv.Atoi(seasonQuery); err != nil {
					return season
				}
			}

			return defaultSeason
	   }

	var leagueID int
	if league == "EPL" {
		leagueID = 44
	} else {
		leagueID = 363
	}

	season := getSeason(2023)
	result, err := hc.apiService.PreviousFixtures(leagueID, season ,from, to)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// save to redis
	

	c.IndentedJSON(http.StatusOK, gin.H{"result": result})
}


