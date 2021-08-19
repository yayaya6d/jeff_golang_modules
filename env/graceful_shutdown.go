package env

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	Name       string
	WorkerFunc func(context.Context) error
}

var signalChannel = make(chan os.Signal, 1)

func GracefulShutdownWithTimeout(timeout time.Duration, workers ...Worker) bool {
	// create a signal channel that wait for shutdown signal

	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	osCall := (<-signalChannel).String()
	fmt.Println("receive system call:", osCall)

	// create context with timeout, and then inject it to Workers
	// the worker that injected context will finish when timeout or work completed.
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	wg := &sync.WaitGroup{}
	for _, worker := range workers {
		wg.Add(1)
		go func(worker Worker) {
			if err := worker.WorkerFunc(ctx); err != nil {
				fmt.Println("error occurred while execute", worker.Name, ", err =", err.Error())
			}
			defer wg.Done()
		}(worker)
	}

	shutdownSuccess := false

	go func() {
		// wait for every worker completed and then call cancel
		wg.Wait()
		cancel()
		shutdownSuccess = true
	}()

	// wait for cancel event or timeout
	<-ctx.Done()
	fmt.Println("Graceful shutdown result:", shutdownSuccess)
	return shutdownSuccess
}
