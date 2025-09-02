package main

import (
	"fmt"
	"log"

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
	fmt.Print(redisClient)

	fixtureRepo := &repository.APIRepo{}

	fixtureUC := usecase.NewFixtureUsecase(fixtureRepo)

	router := router.NewRouter(fixtureUC)
	router.Run()
}
