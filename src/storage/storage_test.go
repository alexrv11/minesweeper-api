package storage

import (
	"github.com/deviget/minesweeper-api/src/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type StorageSuite struct {
	suite.Suite
	keyGame string
	game *models.Game
	kvs *KVS
	duration time.Duration
}

func (suite *StorageSuite) SetupSuite(){
	suite.keyGame = "123456"
	suite.game = models.NewGame(4)
	suite.game.Id = suite.keyGame
	suite.duration = time.Duration(time.Second * 5)
	suite.kvs = NewKVS(suite.duration)
}

func (suite *StorageSuite) TestNewKVS(){
	kvs := NewKVS(time.Duration(time.Second * 5))

	assert.NotNil(suite.T(), kvs)
}

func (suite *StorageSuite) TestPutGame() {

	err := suite.kvs.PutGame(*suite.game)

	assert.Nil(suite.T(), err)
}

func (suite *StorageSuite) TestGetGame() {

	_ = suite.kvs.PutGame(*suite.game)

	result, err := suite.kvs.GetGame(suite.keyGame)

	assert.Equal(suite.T(), suite.keyGame, result.Id)
	assert.Nil(suite.T(), err)
}

func (suite *StorageSuite) TestGetGameExpired() {

	_ = suite.kvs.PutGame(*suite.game)

	time.Sleep(6 * time.Second)

	result, err := suite.kvs.GetGame(suite.keyGame)

	assert.Nil(suite.T(), result)
	assert.NotNil(suite.T(), err)
}

func TestStorageSuite(t *testing.T) {
	suite.Run(t, new(StorageSuite))
}