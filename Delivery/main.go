package main

import (
	"fmt"
	"log"

	"github.com/abrshodin/ethio-fb-backend/Delivery/Router"
	infrastrucutre "github.com/abrshodin/ethio-fb-backend/Infrastructure"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error in loading .env file")
	}
	redisClient := infrastrucutre.RedisConnect()
	fmt.Print(redisClient)
	router := router.NewRouter()
	router.Run()
}
