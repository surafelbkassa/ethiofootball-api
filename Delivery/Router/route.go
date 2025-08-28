package router

import (
	"os"

	controller "github.com/abrshodin/ethio-fb-backend/Delivery/Controller"
	infrastructure "github.com/abrshodin/ethio-fb-backend/Infrastructure"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

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
