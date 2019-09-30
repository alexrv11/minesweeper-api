package domain

import (
	"github.com/deviget/minesweeper-api/src/domain/mocks"
	"github.com/deviget/minesweeper-api/src/models"
	"github.com/deviget/minesweeper-api/src/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type MineSweeperSuite struct {
	suite.Suite
	minesweeper *MineSweeper
	factoryMock *mocks.Factory
	minerMock *mocks.MinerStrategy
}

func (suite *MineSweeperSuite) SetupTest(){
	suite.minerMock = &mocks.MinerStrategy{}
	suite.factoryMock = &mocks.Factory{ MinerMock: suite.minerMock }

	kvs := storage.KVS{ Duration:time.Duration(time.Second * 5), Games: make(map[string]*models.Game, 0)}
	suite.minesweeper = &MineSweeper{ factory: suite.factoryMock, kvs:  kvs}
}

func (suite *MineSweeperSuite) TestCreateGame(){
	dimension := 4
	bombs := 1

	expectedRow := 1
	expectedColumn := 1
	suite.minerMock.On("SelectField").Return(expectedRow, expectedColumn)

	game , err := suite.minesweeper.CreateGame(dimension, bombs)

	isBomb := game.Table[expectedRow][expectedColumn].IsBomb
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), game)
	assert.True(suite.T(), isBomb)
}


func TestMineSweeperSuite(t *testing.T){
	suite.Run(t, new(MineSweeperSuite))
}
