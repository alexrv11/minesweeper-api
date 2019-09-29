package app

import (
	"github.com/gin-gonic/gin"
	"github.com/google/logger"
)

// StartApplication start set up the environment of the application
func StartApplication() *gin.Engine {
	logger.Info("Starting Minesweeper Api")

	router := gin.Default()

	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	LoadResources(router)

	return router
}


