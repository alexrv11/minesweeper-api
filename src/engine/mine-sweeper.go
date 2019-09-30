package engine

import (
	"github.com/deviget/minesweeper-api/src/models"
)

type MineSweeper interface {
	CreateGame(dimension, numberOfBomb int) (*models.Game, error)
	GetGame(id string) (*models.Game, error)
	PlayGame(idGame string, position models.Position) (*models.PlayResult, error)
}
