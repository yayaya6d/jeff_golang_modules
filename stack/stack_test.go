package stack

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackTestSuite struct {
	suite.Suite
	s Stack
}

func (suite *StackTestSuite) SetupTest() {
	suite.s = NewStack()
}

func (suite *StackTestSuite) TestStack_PushSomethingToStack_GetCorrectSize() {
	expectedSize := 10
	for i := 0; i < expectedSize; i++ {
		suite.s.Push(i)
	}

	suite.Equal(expectedSize, suite.s.Size(), "unexpected size, expect %d, actual: %d", expectedSize, suite.s.Size())
}

func (suite *StackTestSuite) TestStack_PushSomethingToStack_GetValFromTopAndGetCorrectVal() {
	size := 10
	for i := 1; i <= size; i++ {
		suite.s.Push(i)
	}

	for i := size; i > 0; i-- {
		cur := suite.s.Top()
		suite.s.Pop()
		suite.Equal(i, cur, "unexpected val, expect %d, actual: %d", i, cur)
	}
}

func (suite *StackTestSuite) TestStack_PushSomethingToStackAndPopAll_IsEmptyIsTrue() {
	size := 10
	for i := 1; i <= size; i++ {
		suite.s.Push(i)
	}

	for i := size; i > 0; i-- {
		suite.s.Pop()
	}

	suite.Equal(true, suite.s.IsEmpty(), "unexpected val from s.IsEmpty(), expect %d, actual: %d", true, suite.s.IsEmpty())
}

func TestStackTestSuite(t *testing.T) {
	suite.Run(t, new(StackTestSuite))
}
