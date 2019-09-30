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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "The id does not exist"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (minesweeper *Minesweeper) HandlerPlayGame(c *gin.Context) {
	idGame := c.Param("id")
	var input models.Position
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := minesweeper.CoreEngine.PlayGame(idGame, input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}
