package main

import (
	"github.com/abrshodin/ethio-fb-backend/Delivery/Router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.RegisterRoute(r)

	r.Run()
}
