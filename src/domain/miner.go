package domain

import (
	"github.com/deviget/minesweeper-api/src/engine"
	"math/rand"
)



type RandomMiner struct {
	dimension int
	activeBomb [][]bool
}

func NewRandomMiner(dimension int) engine.MinerStrategy {
	activeBomb := make([][]bool, dimension)
	return &RandomMiner{ dimension: dimension, activeBomb:activeBomb}
}

func (miner *RandomMiner) SelectField() (int, int) {

	row := rand.Intn(miner.dimension)
	column := rand.Intn(miner.dimension)

	return row, column
}