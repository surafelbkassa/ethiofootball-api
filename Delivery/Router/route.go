package routers

import (
	"os"

	controller "github.com/abrshodin/ethio-fb-backend/Delivery/Controller"
	controlller "github.com/abrshodin/ethio-fb-backend/Delivery/Controllers"
	infrastructure "github.com/abrshodin/ethio-fb-backend/Infrastructure"
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
	newsRouter.GET("/liveScores", newsHandler.GetLiveScores)
	return router
}

func RegisterRoute(router *gin.Engine) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	intentParser := infrastructure.NewAIIntentParser(apiKey)
	intentUsecase := usecase.NewParseIntentUsecase(intentParser)
	intentController := controller.NewIntentController(intentUsecase)
	answerComposer := infrastructure.NewAIAnswerComposer(apiKey)
	answerUseCase := usecase.NewAnswerUseCase(answerComposer)
	answerController := controller.NewAnswerController(answerUseCase)

	router.POST("/intent/parse", intentController.ParseIntent)
	router.POST("/answer", answerController.HandlePostAnswer)
}

func RegisterTeamRoutes(r *gin.Engine, handler *controlller.TeamController) {
	team := r.Group("team")
	{
		team.GET("/:id/bio", handler.GetTeam)
		team.POST("/create", handler.AddTeam)
	}
}

func RegisterAPISercice(r *gin.Engine, handler *controlller.PrevFixturesController) {

	api := r.Group("api")
	{
		api.GET("/previous-fixtures", handler.PreviousMatchHistory)
	}

}
