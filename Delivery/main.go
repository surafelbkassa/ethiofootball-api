package main

import (
	"log"

	controller "github.com/abrshodin/ethio-fb-backend/Delivery/Controller"
	routers "github.com/abrshodin/ethio-fb-backend/Delivery/Router"
	infrastructure "github.com/abrshodin/ethio-fb-backend/Infrastructure"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading .env file")
	}

	// Redis & Team setup
	redisClient := infrastructure.RedisConnect()
	teamRepo := repository.NewTeamRepo(redisClient)
	
	apiService := infrastructure.NewAPIService()
	prevRepo := repository.NewPrevFixturesRepo(redisClient)
	prevUC := usecase.NewFixturesUsecase(apiService, prevRepo)

	teamUsecase := usecase.NewTeamUsecase(teamRepo, apiService)
	teamHandler := controller.NewTeamController(teamUsecase)
	historyHandler := controller.NewFixturesController(prevUC)

	fixtureRepo := repository.NewAPIRepo(redisClient)
	fixtureUC := usecase.NewFixtureUsecase(fixtureRepo, fixtureRepo)

	// News setup
	eventRepo := repository.NewEventRepository()
	newsUC := usecase.NewNewsUseCase(eventRepo)

	// Standings setup
	standingsRepo := repository.NewStandingsRepo(redisClient)
	standingsUC := usecase.NewStandingsUsecase(standingsRepo)
	standingsHandler := controller.NewStandingsController(standingsUC)

	// News route
	newsHandler := controller.NewNewsController(newsUC)
	
	apiKey := os.Getenv("GEMINI_API_KEY")
	answerComposer := infrastructure.NewAIAnswerComposer(apiKey)
	answerUseCase := usecase.NewAnswerUseCase(answerComposer)
	answerController := controller.NewAnswerController(answerUseCase)

	intentParser := infrastructure.NewAIIntentParser(apiKey)
	intentUsecase := usecase.NewParseIntentUsecase(intentParser)
	intentController := controller.NewIntentController(
														intentUsecase,
														standingsHandler,
														newsHandler,
														teamHandler,
														answerController,
													)
	
	// Router
	router := routers.NewRouter(fixtureUC, newsUC)
	routers.RegisterTeamRoutes(router, teamHandler)
	routers.RegisterAPISercice(router, historyHandler)
	routers.RegisterStandingsRoutes(router, standingsHandler)
	routers.RegisterNewsRoutes(router, newsHandler)
	routers.RegisterRoute(router, intentController, answerController)

	router.Run()
}
