package app

import (
	"github.com/gin-gonic/gin"
)

// StartApplication start set up the environment of the application
func StartApplication() *gin.Engine {
	//logger.Info("Starting Minesweeper Api")

	router := gin.New()

	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	LoadResources(router)

	return router
}


