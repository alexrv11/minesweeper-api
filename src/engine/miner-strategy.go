package engine

type MinerStrategy interface {
	SelectField() (int, int)
}