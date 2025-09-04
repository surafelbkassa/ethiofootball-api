package controlller

import (
	"net/http"
	"strconv"

	domain2 "github.com/abrshodin/ethio-fb-backend/Domain"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

type PrevFixturesController struct {
	prevUC usecase.IPrevFixturesUsecase
}

func NewPrevFixturesController(uc usecase.IPrevFixturesUsecase) *PrevFixturesController {
	return &PrevFixturesController{prevUC: uc}
}

func (hc *PrevFixturesController) PreviousMatchHistory(c *gin.Context) {

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
	q := domain2.RoundQuery{League: league, Season: season, Round: round, From: from, To: to}

	rq, err := hc.prevUC.ResolveRoundWindow(c.Request.Context(), q)
	if err == nil {
		q = rq
	}

	cached, err := hc.prevUC.GetCachedByRound(c.Request.Context(), q)
	if err == nil {
		c.IndentedJSON(http.StatusOK, gin.H{"result": cached, "source": "cache"})
		return
	}

	result, err := hc.prevUC.FetchAndStore(c.Request.Context(), league, leagueID, q)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"result": result, "source": "api"})
}
