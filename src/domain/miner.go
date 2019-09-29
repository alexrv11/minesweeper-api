package domain

import "math/rand"


type MinerStrategy interface {
	SelectField() (int, int)
}

type RandomMiner struct {
	dimension int
	activeBomb map[string]bool
}

func NewRandomMiner(dimension int) MinerStrategy {
	activeBomb := make([][]bool, dimension)
	return &RandomMiner{ dimension: dimension, activeBomb:activeBomb}
}

func (miner *RandomMiner) SelectField() (int, int) {

	row := rand.Intn(miner.dimension)
	column := rand.Intn(miner.dimension)

	return row, column
}