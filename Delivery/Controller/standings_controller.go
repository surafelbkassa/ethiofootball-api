package controller

import (
	"net/http"
	"strconv"

	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

type StandingsController struct {
	standingsUsecase usecase.IStandingsUsecase
}

func NewStandingsController(standingsUsecase usecase.IStandingsUsecase) *StandingsController {
	return &StandingsController{
		standingsUsecase: standingsUsecase,
	}
}

func (c *StandingsController) GetStandings(ctx *gin.Context) {
	
	league := ctx.Query("league")
	seasonQuery := ctx.Query("season")

	
	if league == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "league parameter is required"})
		return
	}
	if seasonQuery == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "season parameter is required"})
		return
	}

	leagueID := 0
	if league == "ETH" {
		leagueID = 363
	} else if league == "EPL" {
		leagueID = 39
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "unsupported league"})
		return
	}

	season, err := strconv.Atoi(seasonQuery)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "season must be a valid year between 2021 and 2023"})
		return
	}

	if  season < 2021 || season > 2023 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "season must be a valid year between 2021 and 2023"})
		return
	}

	// Get standings from usecase
	standings, err := c.standingsUsecase.GetStandings(ctx, leagueID, season)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get standings: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, standings)
}


