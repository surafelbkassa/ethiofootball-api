package controller

import (
	"net/http"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	teamUsecase usecase.TeamUsecases
}
func NewTeamController(teamUsecase usecase.TeamUsecases) *TeamController {
	return &TeamController{teamUsecase: teamUsecase}
}

func(tc *TeamController) GetTeam (c *gin.Context) {

	id := c.Param("id")
	ctx := c.Request.Context()

	team, err := tc.teamUsecase.GetTeam(ctx, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": team})
}

func (tc *TeamController) AddTeam(c *gin.Context){
	
	ctx := c.Request.Context()

	var team domain.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid input format"})
		return
	}

	err := tc.teamUsecase.AddTeam(ctx, &team)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "team added successfully"})
}





