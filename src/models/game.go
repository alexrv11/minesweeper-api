package models

import (
	"github.com/google/uuid"
	"time"
)

type Game struct {
	Id          string
	Dimension   int
	Table       [][]*Cell
	CreatedAt   time.Time
	Bombs       int
	HiddenCells int
}


//NewGame creates a new instance of Game
func NewGame(dimension int) *Game {

	id := uuid.New().String()
	table := make([][]*Cell, dimension)
	for row := range table {
		table[row] = make([]*Cell, dimension)
		for column := range table[row] {
			table[row][column] = &Cell{ IsBomb: false, Visited:false, ConnectedBomb: 0 }
		}
	}

	return &Game{ Id: id, Table: table, Dimension: dimension, CreatedAt: time.Now(), Bombs:0, HiddenCells: 0 }
}