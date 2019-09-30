package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GameSuite struct {
	suite.Suite
	game *Game
	dimension int
}

func (suite *GameSuite) SetupSuite(){
	suite.dimension = 4
	game := NewGame(suite.dimension)

	_ = game.ActivateBomb(1,0)
	_ = game.ActivateBomb(1,1)
	_ = game.ActivateBomb(2,0)
	suite.game = game
}

func (suite *GameSuite) TestNewGame() {
	game := NewGame(3)

	assert.NotNil(suite.T(), game)
	assert.Equal(suite.T(), 3, len(game.instance.Table))
}

func (suite *GameSuite) TestActivateBomb(){
	game := NewGame(3)
	//Result nines field
	/*
		221
		**1
		221
	*/
	err1 := game.ActivateBomb(1,0)
	err2 := game.ActivateBomb(1,1)
	instance := game.instance

	assert.Nil(suite.T(), err1)
	assert.Nil(suite.T(), err2)
	assert.Equal(suite.T(), 2, instance.Table[0][0].ConnectedBomb)
	assert.Equal(suite.T(), 2, instance.Table[0][1].ConnectedBomb)
	assert.Equal(suite.T(), 1, instance.Table[0][2].ConnectedBomb)
	assert.True(suite.T(), instance.Table[1][0].IsBomb)
	assert.True(suite.T(), instance.Table[1][1].IsBomb)
	assert.Equal(suite.T(), 1, instance.Table[1][2].ConnectedBomb)
	assert.Equal(suite.T(), 2, instance.Table[2][0].ConnectedBomb)
	assert.Equal(suite.T(), 2, instance.Table[2][1].ConnectedBomb)
	assert.Equal(suite.T(), 1, instance.Table[2][2].ConnectedBomb)
}

func (suite *GameSuite) TestPlayLuck_SelectBomb() {
	/*
		2210
		**10
		*310
		1100
	*/
	result, err := suite.game.PlayLuckInCell(1,1)

	assert.Nil(suite.T(), err)
	assert.True(suite.T(), result.GameOver)
	assert.Equal(suite.T(), 0, len(result.WinPosition))

}

func (suite *GameSuite) TestPlayLuck_SelectACellConnectedBomb() {
	/*
		2210
		**10
		*310
		1100
	*/
	result, err := suite.game.PlayLuckInCell(2,1)

	assert.Nil(suite.T(), err)
	assert.False(suite.T(), result.GameOver)
	assert.Equal(suite.T(), 1, len(result.WinPosition))
	assert.Equal(suite.T(), 2, result.WinPosition[0].Row)
	assert.Equal(suite.T(), 1, result.WinPosition[0].Column)

}

func (suite *GameSuite) TestPlayLuck_SelectSecureCell() {
	/*
		2210
		**10
		*310
		1100
	*/
	result, err := suite.game.PlayLuckInCell(2,3)

	assert.Nil(suite.T(), err)
	assert.False(suite.T(), result.GameOver)
	assert.Equal(suite.T(), 13, len(result.WinPosition))
}

func TestGameSuite(t *testing.T){
	suite.Run(t, new(GameSuite))
}
