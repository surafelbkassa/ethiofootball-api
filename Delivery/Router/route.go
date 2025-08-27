package router

import (
	"os"

	"github.com/abrshodin/ethio-fb-backend/Delivery/Controller"
	"github.com/abrshodin/ethio-fb-backend/Infrastructure"
	"github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

func RegisterRoute(router *gin.Engine) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	intentParser := infrastructure.NewAIIntentParser(apiKey)
	intentUsecase := usecase.NewParseIntentUsecase(intentParser)
	intentController := controller.NewIntentController(intentUsecase)

	router.POST("/intent/parse", intentController.ParseIntent)
}
