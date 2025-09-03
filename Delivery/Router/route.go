package routers

import (
	"net/http"

	controlller "github.com/abrshodin/ethio-fb-backend/Delivery/Controllers"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(fixtureUC usecase.FixtureUsecase, newsUC *usecase.NewsUseCase) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Fixtures route
	router.GET("/fixtures", func(c *gin.Context) {
		league := c.Query("league")
		team := c.Query("team")
		from := c.Query("from")
		to := c.Query("to")

		fixtures, err := fixtureUC.GetFixtures(league, team, from, to)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"fixtures": fixtures})
	})

	// News route 
	newsHandler := controlller.NewNewsController(newsUC)
	newsRouter := router.Group("/news")

	newsRouter.GET("/pastMatches", newsHandler.GetNews)
	newsRouter.GET("/standings", newsHandler.GetStandingNews)
	newsRouter.GET("/futureMatches", newsHandler.GetFutureNews)
	
	return router
}

func RegisterTeamRoutes(r *gin.Engine, handler *controlller.TeamController) {
	team := r.Group("team")
	{
		team.GET("/:id/bio", handler.GetTeam)
		team.POST("/create", handler.AddTeam)
	}
}

func RegisterAPISercice(r *gin.Engine, handler *controlller.HistoryController){

	api := r.Group("api")
	{
		api.GET("/previous-fixtures", handler.PreviousMatchHistory)
	}

}
