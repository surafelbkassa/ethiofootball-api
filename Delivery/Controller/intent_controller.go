package controlller

import (
	"errors"
	"log"
	"net/http"

	"github.com/abrshodin/ethio-fb-backend/Usecase"
	"github.com/gin-gonic/gin"
)

type IntentController struct {
	parseIntent *usecase.ParseIntentUseCase
}

func NewIntentController(parseIntent *usecase.ParseIntentUseCase) *IntentController {
	return &IntentController{parseIntent: parseIntent}
}

func (h *IntentController) ParseIntent(c *gin.Context) {
	var req struct {
		Text string `json:"text"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	intent, err := h.parseIntent.Execute(req.Text)
	if err != nil {
		// map semantic errors to HTTP codes
		switch {
		case errors.Is(err, usecase.ErrInvalidInput):
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			log.Print(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, intent)
}
