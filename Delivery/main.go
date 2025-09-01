package main

import (
	"log"

	controlller "github.com/abrshodin/ethio-fb-backend/Delivery/Controllers"
	router "github.com/abrshodin/ethio-fb-backend/Delivery/Router"
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

	fixtureRepo := &repository.APIRepo{}
	fixtureUC := usecase.NewFixtureUsecase(fixtureRepo)

	router := router.NewRouter(fixtureUC)
	routers.RegisterTeamRoutes(router, teamHandler)

	router.Run()
}
