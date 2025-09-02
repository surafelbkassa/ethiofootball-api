package router

import (
	"net/http"

	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(fixtureUC usecase.FixtureUsecase) *gin.Engine {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/fixtures", func(c *gin.Context) {
		league := c.Query("league")
		team := c.Query("team")
		from := c.Query("from")
		to := c.Query("to")

		// call usecase
		fixtures, err := fixtureUC.GetFixtures(league, team, from, to)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"fixtures": fixtures})
	})

	return router
}
