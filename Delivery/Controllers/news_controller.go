package controlller

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
