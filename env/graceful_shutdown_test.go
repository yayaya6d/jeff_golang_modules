package env

import (
	"context"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type gracefulShutdownTestL1 struct {
	suite.Suite
}

func (s *gracefulShutdownTestL1) TestGracefulShutdown_ShutDownAllWorkersSuccessWithoutTimeout_ReturnTrue() {
	workerExecuteTime := time.Second
	shutdownTimeout := time.Second * 2

	worker1 := &Worker{
		Name: "worker1",
		WorkerFunc: func(c context.Context) error {
			time.Sleep(workerExecuteTime)
			return nil
		},
	}

	worker2 := &Worker{
		Name: "worker2",
		WorkerFunc: func(c context.Context) error {
			time.Sleep(workerExecuteTime)
			return nil
		},
	}

	success := true
	go func() {
		success = GracefulShutdownWithTimeout(shutdownTimeout, *worker1, *worker2)
	}()
	signalChannel <- syscall.SIGINT
	time.Sleep(time.Second * 3)

	s.True(success)
}

func (s *gracefulShutdownTestL1) TestGracefulShutdown_OneOfWorkersTimeout_ReturnFalse() {
	shutdownTimeout := time.Second * 2

	worker1 := &Worker{
		Name: "worker1",
		WorkerFunc: func(c context.Context) error {
			time.Sleep(time.Second)
			return nil
		},
	}

	timeoutWorker := &Worker{
		Name: "timeoutWorker",
		WorkerFunc: func(c context.Context) error {
			time.Sleep(shutdownTimeout + time.Second)
			return nil
		},
	}

	success := true
	go func() {
		success = GracefulShutdownWithTimeout(shutdownTimeout, *worker1, *timeoutWorker)
	}()
	signalChannel <- syscall.SIGINT
	time.Sleep(time.Second * 4)

	s.False(success)
}

func TestGracefulShutdownSuite(t *testing.T) {
	suite.Run(t, new(gracefulShutdownTestL1))
}
