package domain

import "github.com/deviget/minesweeper-api/src/engine"

//Factory implements the builds methods
type Factory struct {

}

//BuildMiner builds a miner strategy.
func (factory *Factory) BuildMiner(dimension int) engine.MinerStrategy {
	return NewRandomMiner(dimension)
}