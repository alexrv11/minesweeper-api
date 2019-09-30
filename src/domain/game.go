package domain

import (
	"fmt"
	"github.com/deviget/minesweeper-api/src/models"
	"github.com/pkg/errors"
)

type Game struct {
	instance *models.Game
}

func NewGame(dimension int) *Game{
	return &Game{instance: models.NewGame(dimension)}
}

func (game *Game) ActivateBomb(row, column int) error {

	err := game.validatePosition(row, column)

	if err != nil {
		return err
	}

	game.instance.Table[row][column].IsBomb = true
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
	if row >= 0 && row < game.instance.Dimension && column >=0 && column < game.instance.Dimension {
		cell := game.instance.Table[row][column]
		cell.ConnectedBomb += 1
	}
}

func (game *Game) PlayLuckInCell(row, column int) (*models.PlayResult, error) {

	err := game.validatePosition(row, column)

	if err != nil {
		return nil, err
	}

	winPosition := make([]models.Position, 0)
	result := &models.PlayResult{WinPosition: winPosition, GameOver: false }

	cell := game.instance.Table[row][column]
	if cell.IsBomb {
		result.GameOver = true
		return result, nil
	}

	if cell.ConnectedBomb > 0 {
		position := models.Position{Row: row, Column:column}
		result.WinPosition = append(result.WinPosition, position)
		return result, nil
	}

	//the cell is empty, So we are going to list all the neighbor cell that are not bombs
	game.expandSecureCells(row, column, result)

	return result, nil
}

func (game *Game) expandSecureCells(row, column int, playResult *models.PlayResult) {

	err := game.validatePosition(row, column)

	//when the position is invalid don't make anything
	if err != nil {
		return
	}

	cell := game.instance.Table[row][column]

	if cell.Visited || cell.IsBomb {
		return
	}
	cell.Visited = true
	playResult.WinPosition = append(playResult.WinPosition, models.Position{Row:row, Column:column})

	//propagate the secure cells
	game.expandSecureCells(row - 1, column, playResult)
	game.expandSecureCells(row, column + 1, playResult)
	game.expandSecureCells(row + 1, column + 1, playResult)
	game.expandSecureCells(row, column - 1, playResult)
}

func (game *Game) validatePosition(row, column int) error {
	if row >= game.instance.Dimension || row < 0 {
		return errors.New(fmt.Sprintf("The row %d is out of the range in the game dimesion %d", row, game.instance.Dimension))
	}

	if column >= game.instance.Dimension || column < 0 {
		return errors.New(fmt.Sprintf("The column %d is out of the range in the game dimesion %d", row, game.instance.Dimension))
	}

	return nil
}