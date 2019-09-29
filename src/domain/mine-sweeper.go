package domain

import "github.com/deviget/minesweeper-api/src/models"

type MineSweeperEngine interface {
	CreateGame(dimension, numberOfBomb int) (*models.Game, error)
}

type MineSweeper struct { factory FactoryEngine}

//CreateGame creates a Game instance given dimension of the table n x n and the number of bomb that is going to be added in the table
func (mineSweeper *MineSweeper) CreateGame(dimension, numberOfBomb int) (*models.Game, error) {
	game := models.NewGame(dimension)

	//Mining bombs
	counterMining := 0
	miner := mineSweeper.factory.BuildMiner()
	for counterMining < numberOfBomb {
		row, column := miner.SelectField(dimension)
		err := game.ActivateBomb(row, column)
		if err != nil {
			return nil, err
		}

		counterMining++
	}


	return game
}
