package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FactorySuite struct {
	suite.Suite
	factory Factory
}

func (suite *FactorySuite) TestBuildMiner(){
	miner := suite.factory.BuildMiner(3)

	assert.NotNil(suite.T(), miner)
}

func TestFactorySuite(t *testing.T){
	suite.Run(t, new(FactorySuite))
}
