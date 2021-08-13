package queue

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ListQueueTestSuite struct {
	suite.Suite
	q ListQueue
}

func (suite *ListQueueTestSuite) SetupTest() {
	suite.q = NewListQueue()
}

func (suite *ListQueueTestSuite) TestSize_PushSomethingIntoListQueueAndPopSometimes_GetCorrectSize() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	for i := 0; i < 3; i++ {
		suite.q.Pop()
	}

	suite.Equal(7, suite.q.Size(), "expected size is %d, got %d", 0, suite.q.Size())
}

func (suite *ListQueueTestSuite) TestIsEmpty_PushSomethingIntoListQueue_GetIsNotEmpty() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	suite.Equal(false, suite.q.IsEmpty(), "expect queue is not empty but is.")
}

func (suite *ListQueueTestSuite) TestIsEmpty_PushSomethingIntoListQueueAndPopAll_GetIsEmpty() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	for i := 0; i < 10; i++ {
		suite.q.Pop()
	}

	suite.Equal(true, suite.q.IsEmpty(), "expect queue is empty but not.")
}

func (suite *ListQueueTestSuite) TestQueue_PushSomethingIntoListQueue_GetCorrectFrontValAndBackVal() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	suite.Equal(0, suite.q.Front(), "expected front val is %d, got %d", 0, suite.q.Front())
	suite.Equal(9, suite.q.Back(), "expected back val is %d, got %d", 0, suite.q.Back())
}

func (suite *ListQueueTestSuite) TestQueue_PushSomethingIntoListQueueAndPopSometimes_GetCorrectFrontValAndBackVal() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	for i := 0; i < 3; i++ {
		suite.q.Pop()
	}

	suite.Equal(3, suite.q.Front(), "expected front val is %d, got %d", 0, suite.q.Front())
	suite.Equal(9, suite.q.Back(), "expected back val is %d, got %d", 0, suite.q.Back())
}

func TestListQueueTestSuite(t *testing.T) {
	suite.Run(t, new(ListQueueTestSuite))
}
