package main

import (
	"log"

	controller "github.com/abrshodin/ethio-fb-backend/Delivery/Controllers"
	routers "github.com/abrshodin/ethio-fb-backend/Delivery/Router"
	infrastructure "github.com/abrshodin/ethio-fb-backend/Infrastructure"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error in loading .env file")
	}

	// Redis & Team setup
	redisClient := infrastructure.RedisConnect()
	teamRepo := repository.NewTeamRepo(redisClient)
	teamUsecase := usecase.NewTeamUsecase(teamRepo)
	teamHandler := controller.NewTeamController(teamUsecase)

	apiService := infrastructure.NewAPIService()
	prevRepo := repository.NewPrevFixturesRepo(redisClient)
	prevUC := usecase.NewFixturesUsecase(apiService, prevRepo)
	historyHandler := controller.NewFixturesController(prevUC)

	fixtureRepo := &repository.APIRepo{}
	fixtureUC := usecase.NewFixtureUsecase(fixtureRepo)

	// News setup
	eventRepo := repository.NewEventRepository()
	newsUC := usecase.NewNewsUseCase(eventRepo)

	// Router
	router := routers.NewRouter(fixtureUC, newsUC)
	routers.RegisterTeamRoutes(router, teamHandler)
	routers.RegisterAPISercice(router, historyHandler)

	router.Run()
}
