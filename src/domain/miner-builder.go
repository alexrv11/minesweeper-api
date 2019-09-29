package domain

type FactoryEngine interface {
	BuildMiner(dimension int) MinerStrategy
}

type Factory interface {
}

func (factory *Factory) BuildMiner(dimension int) MinerStrategy {
	return NewRandomMiner(dimension)
}