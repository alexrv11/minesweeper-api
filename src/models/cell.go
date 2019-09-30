package models

type Cell struct {
	IsBomb bool
	ConnectedBomb int
	Visited bool
}
