package domain

import (
	"github.com/deviget/minesweeper-api/src/domain/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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
	suite.minesweeper = &MineSweeper{ factory: suite.factoryMock }
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
