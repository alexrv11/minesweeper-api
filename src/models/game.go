package models

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

type Game struct {
	Id string
	Dimension int
	Table [][]*Cell
	CreatedAt time.Time
	Timer *time.Timer
}

//NewGame creates a new instance of Game
func NewGame(dimension int) *Game {

	id := uuid.New().String()
	table := make([][]*Cell, dimension)
	for row := range table {
		table[row] = make([]*Cell, dimension)
		for column := range table[row] {
			table[row][column] = &Cell{ IsBomb: false, ConnectedBomb: 0 }
		}
	}

	return &Game{ Id: id, Table: table, Dimension: dimension, CreatedAt: time.Now() }
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

func (game *Game) NotifyNeighborCell(row, column int) {
	game.incrementConnectedBomb(row - 1, column)
	game.incrementConnectedBomb(row - 1, column + 1)
	game.incrementConnectedBomb(row, column + 1)
	game.incrementConnectedBomb(row + 1, column + 1)
	game.incrementConnectedBomb(row + 1, column)
	game.incrementConnectedBomb(row + 1, column - 1)
	game.incrementConnectedBomb(row, column - 1)
	game.incrementConnectedBomb(row-1, column - 1)
}

func (game *Game) incrementConnectedBomb(row, column int) {
	if row >= 0 && row < game.Dimension && column >=0 && column < game.Dimension {
		cell := game.Table[row][column]
		cell.ConnectedBomb += 1
	}
}
