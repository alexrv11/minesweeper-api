package domain

import (
	"github.com/deviget/minesweeper-api/src/engine"
	"github.com/deviget/minesweeper-api/src/models"
	"github.com/deviget/minesweeper-api/src/storage"
)


type MineSweeper struct {
	factory engine.Factory
	kvs storage.KVS
}

func NewMinesweeper(factory engine.Factory, kvs storage.KVS) engine.MineSweeper {
	return &MineSweeper{ factory:factory, kvs:kvs }
}

//CreateGame creates a Game instance given dimension of the table n x n and the number of bomb that is going to be added in the table
func (mineSweeper *MineSweeper) CreateGame(dimension, numberOfBomb int) (*models.Game, error) {
	game := models.NewGame(dimension)

	//Mining bombs
	counterMining := 0
	miner := mineSweeper.factory.BuildMiner(dimension)
	for counterMining < numberOfBomb {
		row, column := miner.SelectField()
		err := game.ActivateBomb(row, column)
		if err != nil {
			return nil, err
		}

		counterMining++
	}

	_ := mineSweeper.kvs.PutGame(*game)

	return game, nil
}



