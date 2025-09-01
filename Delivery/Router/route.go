package routers

import (
	"net/http"

	controlller "github.com/abrshodin/ethio-fb-backend/Delivery/Controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}

func RegisterTeamRoutes(r *gin.Engine, handler *controlller.TeamController){

	team := r.Group("team")
	{
		team.GET("/:id/bio", handler.GetTeam)
		team.POST("/create", handler.AddTeam)
	}
}
