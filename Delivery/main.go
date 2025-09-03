package main

import (
	"log"

	controlller "github.com/abrshodin/ethio-fb-backend/Delivery/Controllers"
	routers "github.com/abrshodin/ethio-fb-backend/Delivery/Router"
	infrastrucutre "github.com/abrshodin/ethio-fb-backend/Infrastructure"
	repository "github.com/abrshodin/ethio-fb-backend/Repository"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading .env file")
	}

	redisClient := infrastrucutre.RedisConnect()
	teamRepo := repository.NewTeamRepo(redisClient)
	teamUsecase := usecase.NewTeamUsecase(teamRepo)
	teamHandler := controlller.NewTeamController(teamUsecase)

	HistoryService := infrastrucutre.NewHistoryService()
	historyHandler := controlller.NewHistoryController(HistoryService)

	fixtureRepo := &repository.APIRepo{}
	fixtureUC := usecase.NewFixtureUsecase(fixtureRepo)

	router := routers.NewRouter(fixtureUC)
	routers.RegisterTeamRoutes(router, teamHandler)
	routers.RegisterAPISercice(router, historyHandler)

	router.Run()
}
