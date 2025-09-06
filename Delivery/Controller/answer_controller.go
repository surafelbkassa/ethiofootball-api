package controller

import (
	"net/http"
	"time"

	domain "github.com/abrshodin/ethio-fb-backend/Domain"
	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

type AnswerController struct {
	answerUsecase usecase.AnswerUsecase
}

func NewAnswerController(auc usecase.AnswerUsecase) *AnswerController {
	return &AnswerController{
		answerUsecase: auc,
	}
}

type postAnswerRequest struct {
	Topic       string                 `json:"topic"`
	Language    string                 `json:"language"`
	Source      string                 `json:"source"`
	Freshness   time.Time              `json:"freshness"`
	ContextData map[string]interface{} `json:"context_data"`
}

func (c *AnswerController) HandlePostAnswer(ctx *gin.Context) {
	var req postAnswerRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		return
	}

	answerContext := domain.AnswerContext{
		Topic:       req.Topic,
		Language:    req.Language,
		Source:      req.Source,
		Freshness:   req.Freshness,
		ContextData: req.ContextData,
	}

	answer, err := c.answerUsecase.Compose(ctx, answerContext)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to compose answer"})
		return
	}

	ctx.JSON(http.StatusOK, answer)
}
