package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

// StartApplication start set up the environment of the application
func StartApplication() *gin.Engine {
	logger.Info("Starting Minesweeper Api")

	router := gin.Default()

	LoadResources(router)

	return router
}


