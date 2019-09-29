package models

import (
	"fmt"
	"github.com/pkg/errors"
)

type Game struct {
	Dimension int
	Table [][]*Cell
}

//NewGame creates a new instance of Game
func NewGame(dimension int) *Game {

	table := make([][]*Cell, dimension)
	for row := range table {
		table[row] = make([]*Cell, dimension)
	}

	return &Game{ Table: table, Dimension: dimension }
}

func (game *Game) ActivateBomb(row, column int) error {
	if row >= game.Dimension {
		return errors.New(fmt.Sprintf("The row %d is out of the range in the game dimesion %d", row, game.Dimension))
	}

	if column >= game.Dimension {
		return errors.New(fmt.Sprintf("The column %d is out of the range in the game dimesion %d", row, game.Dimension))
	}

	game.Table[row][column].IsBomb = true

	return nil
}