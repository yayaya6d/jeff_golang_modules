package queue

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type QueueTestSuite struct {
	suite.Suite
	q Queue
}

func (suite *QueueTestSuite) SetupTest() {
	suite.q = NewQueue()
}

func (suite *QueueTestSuite) TestSize_PushSomethingIntoQueueAndPopSometimes_GetCorrectSize() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	for i := 0; i < 3; i++ {
		suite.q.Pop()
	}

	suite.Equal(7, suite.q.Size(), "expected size is %d, got %d", 0, suite.q.Size())
}

func (suite *QueueTestSuite) TestIsEmpty_PushSomethingIntoQueue_GetIsNotEmpty() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	suite.Equal(false, suite.q.IsEmpty(), "expect queue is not empty but is.")
}

func (suite *QueueTestSuite) TestIsEmpty_PushSomethingIntoQueueAndPopAll_GetIsEmpty() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	for i := 0; i < 10; i++ {
		suite.q.Pop()
	}

	suite.Equal(true, suite.q.IsEmpty(), "expect queue is empty but not.")
}

func (suite *QueueTestSuite) TestQueue_PushSomethingIntoQueue_GetCorrectFrontValAndBackVal() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	suite.Equal(0, suite.q.Front(), "expected front val is %d, got %d", 0, suite.q.Front())
	suite.Equal(9, suite.q.Back(), "expected back val is %d, got %d", 0, suite.q.Back())
}

func (suite *QueueTestSuite) TestQueue_PushSomethingIntoQueueAndPopSometimes_GetCorrectFrontValAndBackVal() {
	for i := 0; i < 10; i++ {
		suite.q.Push(i)
	}

	for i := 0; i < 3; i++ {
		suite.q.Pop()
	}

	suite.Equal(3, suite.q.Front(), "expected front val is %d, got %d", 0, suite.q.Front())
	suite.Equal(9, suite.q.Back(), "expected back val is %d, got %d", 0, suite.q.Back())
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}
