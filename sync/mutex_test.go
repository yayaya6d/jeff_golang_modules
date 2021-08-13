package sync

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type mutexTestSuite struct {
	suite.Suite
	m Mutex
}

func (suite *mutexTestSuite) SetupTest() {
	suite.m = NewMutex()
}

func (s *mutexTestSuite) TestMutex_TestLockAndUnLock_ExecuteCorrectly() {
	// arrange
	count := 0
	testFunc := func() {
		s.m.Lock()
		count++
		s.m.Unlock()
	}

	// act
	go testFunc()
	go testFunc()
	time.Sleep(time.Second)

	// assert
	s.Equal(2, count)
}

func (s *mutexTestSuite) TestMutex_TestTryLock_ExecuteCorrectly() {
	// arrange
	s.m.Lock()

	// act and assert
	s.False(s.m.TryLock())
	s.m.Unlock()
	s.True(s.m.TryLock())
	s.False(s.m.TryLock())
	s.m.Unlock()
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(mutexTestSuite))
}
