package storage

import (
	"github.com/deviget/minesweeper-api/src/models"
	"time"
)

type KVS struct {
	Games map[string]*models.Game
	Duration time.Duration
}

func NewKVS(cleanInterval time.Duration) *KVS {
	return &KVS{Duration: cleanInterval, Games: make(map[string]*models.Game)}
}

func (kvs *KVS) PutGame(game models.Game) error {

	kvs.Games[game.Id] = &game


	return nil
}

func (kvs *KVS) GetGame(id string) (*models.Game, error) {
	game := kvs.Games[id]


	return game, nil
}

func (kvs *KVS) PauseGame(id string) bool {
	//game := kvs.Games[id]

	return true
}
