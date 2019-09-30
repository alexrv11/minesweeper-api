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

func (minesweeper *Minesweeper) HandlerCreateGame(c *gin.Context) {
	var input models.InputGame
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game, _ := minesweeper.CoreEngine.CreateGame(input.Dimension, input.NumberOfBomb)

	c.JSON(http.StatusCreated, game)
}

func (minesweeper *Minesweeper) HandlerGetGame(c *gin.Context){
	id := c.Param("id")
	result, err := minesweeper.CoreEngine.GetGame(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, err) //TODO: improve this error handler, alwasy return BadDRequest, when the id doesn't exist return 404.
	}

	c.JSON(http.StatusOK, result)
}
