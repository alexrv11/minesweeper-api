// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import "github.com/deviget/minesweeper-api/src/engine"

// FactoryEngine is an autogenerated mock type for the FactoryEngine type
type Factory struct {
	MinerMock engine.MinerStrategy
}

// BuildMiner provides a mock function with given fields: dimension
func (factory *Factory) BuildMiner(dimension int) engine.MinerStrategy {

	return factory.MinerMock
}