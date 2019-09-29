package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MinerSuite struct {
	suite.Suite
}

func (suite *MinerSuite) TestNewRandomMiner(){

	miner := NewRandomMiner(4)

	assert.NotNil(suite.T(), miner)
}

func (suite *MinerSuite) TestSelectField() {
	dimension := 3
	miner := NewRandomMiner(dimension)

	row, column := miner.SelectField()

	assert.True(suite.T(), row >= 0)
	assert.True(suite.T(), column >= 0)
	assert.True(suite.T(), row < dimension)
	assert.True(suite.T(), column < dimension)
}

func TestMinerSuite(t *testing.T) {
	suite.Run(t, new(MinerSuite))
}
