package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GameSuite struct {
	suite.Suite
}

func (suite *GameSuite) TestNewGame() {
	game := NewGame(3)

	assert.NotNil(suite.T(), game)
	assert.Equal(suite.T(), 3, len(game.Table))
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

	assert.Equal(suite.T(), 2, game.Table[0][0])
	assert.Equal(suite.T(), 2, game.Table[0][1])
	assert.Equal(suite.T(), 1, game.Table[0][2])
	assert.True(suite.T(), game.Table[1][0].IsBomb)
	assert.True(suite.T(), game.Table[1][1].IsBomb)
	assert.Equal(suite.T(), 1, game.Table[1][2])
	assert.Equal(suite.T(), 2, game.Table[2][0])
	assert.Equal(suite.T(), 2, game.Table[2][1])
	assert.Equal(suite.T(), 2, game.Table[2][2])
}

func TestGameSuite(t *testing.T){
	suite.Run(t, new(GameSuite))
}
