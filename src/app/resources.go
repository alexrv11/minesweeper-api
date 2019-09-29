package app

import (
	"github.com/deviget/minesweeper-api/src/domain"
	"github.com/deviget/minesweeper-api/src/handlers"
	"github.com/gin-gonic/gin"
)

func LoadResources(router *gin.Engine) {

	router.GET("/ping", handlers.HandlerPing)

	factory := &domain.Factory{}
	gameHandler := handlers.NewMinesweeper(domain.NewMinesweeper(factory))
	router.POST("/games", gameHandler.HandlerMinesweeper)
}
