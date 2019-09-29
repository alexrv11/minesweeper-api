package storage

import (
	"github.com/deviget/minesweeper-api/src/models"
	"github.com/pkg/errors"
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
	game.Timer = time.NewTimer(kvs.Duration)
	kvs.Games[game.Id] = &game


	return nil
}

func (kvs *KVS) GetGame(id string) (*models.Game, error) {
	game := kvs.Games[id]

	if Expired(game.Timer) {
		return nil, errors.New("The game has expired")
	}

	return game, nil
}

func (kvs *KVS) PauseGame(id string) bool {
	game := kvs.Games[id]

	return game.Timer.Stop()
}

//Expired verifies the time has expired
func Expired(T *time.Timer) bool {
	select {
	case <-T.C:
		return true
	default:
		return false
	}
}
