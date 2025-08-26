package main

import (
	"github.com/abrshodin/ethio-fb-backend/Delivery/Router"
)

func main() {
	router := router.NewRouter()

	router.Run()
}
