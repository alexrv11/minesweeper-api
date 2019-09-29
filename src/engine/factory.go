package engine

type Factory interface {
	BuildMiner(dimension int) MinerStrategy
}

