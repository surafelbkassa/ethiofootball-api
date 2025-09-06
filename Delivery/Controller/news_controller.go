package controller

import (
	"net/http"

	usecase "github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

type NewsController struct {
	newsUC *usecase.NewsUseCase
}

func NewNewsController(newsUC *usecase.NewsUseCase) *NewsController {
	return &NewsController{newsUC: newsUC}
}

func (c *NewsController) GetNews(ctx *gin.Context) {
	news, err := c.newsUC.GenerateNews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"news":   news,
	})
}

func (c *NewsController) GetStandingNews(ctx *gin.Context) {
	news, err := c.newsUC.GenerateStandingNews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"news":   news,
	})
}

func (c *NewsController) GetFutureNews(ctx *gin.Context) {
	news, err := c.newsUC.GenerateFutureNews()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"news":   news,
	})
}

func (c *NewsController) GetLiveScores(ctx *gin.Context) {
	news, err := c.newsUC.GenerateLiveScores()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"message": news,
	})
}