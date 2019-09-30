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

func (suite *GameSuite) SetupTest(){
	suite.dimension = 4
	game := NewGame(suite.dimension)
	_ = game.ActivateBomb(1,0)
	_ = game.ActivateBomb(1,1)
	_ = game.ActivateBomb(2,0)
	game.instance.HiddenCells = suite.dimension * suite.dimension - 3 //don't count the bombs
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
	assert.Equal(suite.T(), GameOverStatus, result.StatusGame)
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
	assert.Equal(suite.T(), PlayingStatus, result.StatusGame)
	assert.Equal(suite.T(), 1, len(result.WinPosition))
	assert.Equal(suite.T(), 2, result.WinPosition[0].Row)
	assert.Equal(suite.T(), 1, result.WinPosition[0].Column)
}

func (suite *GameSuite) TestPlayLuck_WinTheGame() {
	/*
		XX10
		**10
		*310
		X100
	*/
	expectedHiddenCell := 3
	result, err := suite.game.PlayLuckInCell(2,3)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), PlayingStatus, result.StatusGame)
	assert.Equal(suite.T(), 10, len(result.WinPosition))
	assert.Equal(suite.T(), expectedHiddenCell, suite.game.instance.HiddenCells)
	assert.False(suite.T(), suite.game.instance.Table[0][0].Visited)
	assert.False(suite.T(), suite.game.instance.Table[0][1].Visited)
	assert.False(suite.T(), suite.game.instance.Table[3][0].Visited)
}

func (suite *GameSuite) TestPlayLuck_SelectSecureCell() {
	/*
		XX10
		**10
		*310
		X100
	*/

	result1, err := suite.game.PlayLuckInCell(2,3)
	result2, err := suite.game.PlayLuckInCell(0,1)
	result3, err := suite.game.PlayLuckInCell(0,0)
	result4, err := suite.game.PlayLuckInCell(3,0)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), PlayingStatus, result1.StatusGame)
	assert.Equal(suite.T(), 10, len(result1.WinPosition))
	assert.Equal(suite.T(), PlayingStatus, result2.StatusGame)
	assert.Equal(suite.T(), 1, len(result2.WinPosition))
	assert.Equal(suite.T(), PlayingStatus, result3.StatusGame)
	assert.Equal(suite.T(), 1, len(result3.WinPosition))
	assert.Equal(suite.T(), WinStatus, result4.StatusGame)
	assert.Equal(suite.T(), 1, len(result4.WinPosition))
}

func TestGameSuite(t *testing.T){
	suite.Run(t, new(GameSuite))
}
