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

	err := game.validatePosition(row, column)

	if err != nil {
		return err
	}

	game.Table[row][column].IsBomb = true
	game.NotifyNeighborCell(row, column)

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

func (game *Game) PlayLuckInCell(row, column int) (*PlayResult, error) {

	err := game.validatePosition(row, column)

	if err != nil {
		return nil, err
	}

	var winPosition := make([]Position, 0)
	result := &PlayResult{WinPosition: winPosition, GameOver: false }

	cell := game.Table[row][column]
	if cell.IsBomb {
		result.GameOver = true
		return result, nil
	}

	if cell.ConnectedBomb > 0 {
		position := Position{Row: row, Column:column}
		result.WinPosition = append(result.WinPosition, position)
		return result, nil
	}

	//the cell is empty, So we are going to list all the neighbor cell that are not bombs
	game.expandSecureCells(row, column, result)

	return result, nil
}

func (game *Game) expandSecureCells(row, column int, playResult *PlayResult) {

	err := game.validatePosition(row, column)

	//when the position is invalid don't make anything
	if err != nil {
		return
	}

	playResult.WinPosition = append(playResult.WinPosition, Position{row, column})

	//propagate the secure cells
	game.expandSecureCells(row - 1, column, playResult)
	game.expandSecureCells(row, column + 1, playResult)
	game.expandSecureCells(row + 1, column + 1, playResult)
	game.expandSecureCells(row, column - 1, playResult)

}

func (game *Game) validatePosition(row, column int) error {
	if row >= game.Dimension {
		return errors.New(fmt.Sprintf("The row %d is out of the range in the game dimesion %d", row, game.Dimension))
	}

	if column >= game.Dimension {
		return errors.New(fmt.Sprintf("The column %d is out of the range in the game dimesion %d", row, game.Dimension))
	}

	return nil
}