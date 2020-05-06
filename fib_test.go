package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type fibTestSuite struct {
	suite.Suite
}

func Test_fibTestSuite(t *testing.T) {
	suite.Run(t, new(fibTestSuite))
}

func (suite *fibTestSuite) Test_fib() {
	suite.Assert().Equal(uint64(8), fib(6))
}
