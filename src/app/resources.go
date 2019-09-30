package app

import (
	"github.com/deviget/minesweeper-api/src/domain"
	"github.com/deviget/minesweeper-api/src/handlers"
	"github.com/deviget/minesweeper-api/src/models"
	"github.com/deviget/minesweeper-api/src/storage"
	"github.com/gin-gonic/gin"
	"time"
)

func LoadResources(router *gin.Engine) {

	router.GET("/ping", handlers.HandlerPing)

	factory := &domain.Factory{}
	kvs := storage.KVS{Duration: time.Duration(time.Second*300), Games: map[string]*models.Game{}}
	gameHandler := handlers.NewMinesweeper(domain.NewMinesweeper(factory, kvs))
	router.POST("/games", gameHandler.HandlerMinesweeper)
}
