package handlers

import (
	"github.com/deviget/minesweeper-api/src/engine"
	"github.com/deviget/minesweeper-api/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Minesweeper struct {
	CoreEngine engine.MineSweeper
}

func NewMinesweeper(coreEngine engine.MineSweeper) *Minesweeper {
	return &Minesweeper{ CoreEngine: coreEngine }
}

func (minesweeper *Minesweeper) HandlerMinesweeper(c *gin.Context) {
	var input models.InputGame
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, _ := minesweeper.CoreEngine.CreateGame(input.Dimension, input.NumberOfBomb)

	c.JSON(http.StatusCreated, game)
}
