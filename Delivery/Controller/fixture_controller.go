package controller

import (
	"net/http"
	"strconv"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

type FixturesController struct {
	FixureUC usecase.IFixturesUsecase
}

func NewFixturesController(uc usecase.IFixturesUsecase) *FixturesController {
	return &FixturesController{FixureUC: uc}
}

func (hc *FixturesController) PreviousMatchHistory(c *gin.Context) {

	league := c.Query("league")
	from := c.Query("from")
	to := c.Query("to")
	seasonQuery := c.Query("season")
	round := c.Query("round")

	if league == "" || round == "" {
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
			if season, err := strconv.Atoi(seasonQuery); err == nil {
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
	q := domain.RoundQuery{League: league, Season: season, Round: round, From: from, To: to}

	rq, err := hc.FixureUC.ResolveRoundWindow(c.Request.Context(), q)
	if err == nil {
		q = rq
	}

	cached, err := hc.FixureUC.GetCachedByRound(c.Request.Context(), q)
	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"result": cached, "source": "cache"})
		return
	}

	result, err := hc.FixureUC.FetchAndStore(c.Request.Context(), league, leagueID, q)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": result, "source": "api"})
}

func(fc *FixturesController) LiveFixtures (c *gin.Context){

	league := c.Query("league")
	if league != "EPL" && league != "ETH" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "unsupported league queries"})
		c.Abort()
		return
	}

	result, err := fc.FixureUC.GetLiveMatches(league)
	// if result == nil && err != nil {
	// 	c.IndentedJSON(http.StatusOK, gin.H{"result": result})
	// 	return
	// }

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": result})
}
