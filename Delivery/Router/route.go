package routers

import (
	"net/http"

	controller "github.com/abrshodin/ethio-fb-backend/Delivery/Controller"
	domain "github.com/abrshodin/ethio-fb-backend/Domain"
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
		season := c.Query("season")
		from := c.Query("from")
		to := c.Query("to")

		fixtures, err := fixtureUC.GetFixtures(
			c.Request.Context(), // Pass context
			league,
			team,
			season,
			from,
			to,
		)
		if err != nil {
			// Log error server-side and continue returning empty fixtures
			c.JSON(http.StatusOK, gin.H{"fixtures": []domain.Fixture{}})
			return
		}
		if fixtures == nil {
			fixtures = []domain.Fixture{}
		}
		c.JSON(http.StatusOK, gin.H{"fixtures": fixtures})
	})

	
	
	return router
}

func RegisterNewsRoutes(router *gin.Engine, newsHandler *controller.NewsController){

	newsRouter := router.Group("/news")

	newsRouter.GET("/pastMatches", newsHandler.GetNews)
	newsRouter.GET("/standings", newsHandler.GetStandingNews)
	newsRouter.GET("/futureMatches", newsHandler.GetFutureNews)
	newsRouter.GET("/liveScores", newsHandler.GetLiveScores)
}

func RegisterRoute(router *gin.Engine, handler *controller.IntentController, answerHandler *controller.AnswerController) {

	router.POST("/intent/parse", handler.ParseIntent)
	router.POST("/answer", answerHandler.HandlePostAnswer)
}

func RegisterTeamRoutes(r *gin.Engine, handler *controller.TeamController) {
	team := r.Group("team")
	{
		team.GET("/:id/bio", handler.GetTeam)
		team.POST("/create", handler.AddTeam)
	}
}

func RegisterAPISercice(r *gin.Engine, handler *controller.FixturesController) {

	api := r.Group("api")
	{
		api.GET("/previous-fixtures", handler.PreviousMatchHistory)
		api.GET("/live", handler.LiveFixtures)
	}

}

func RegisterStandingsRoutes(r *gin.Engine, handler *controller.StandingsController) {
	standings := r.Group("api/standings")
	{
		standings.GET("", handler.GetStandings)
	}
}
