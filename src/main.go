package main

import (
	"fmt"
	"github.com/deviget/minesweeper-api/src/app"
	"github.com/deviget/minesweeper-api/src/config"
	"github.com/gin-gonic/gin"
)

func main() {

	configuration := config.LoadConfig()

	gin.SetMode(gin.DebugMode)

	router := app.StartApplication()

	_ = router.Run(fmt.Sprintf(":%s", configuration.Port))
}
