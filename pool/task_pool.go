package pool

import "sync"

type TaskPool[T any] interface {
	Start()
	Submit(task T)
	Close()
}

type taskPool[T any] struct {
	taskQueue   chan T
	taskFunc    func(T)
	workerCount int
	wg          sync.WaitGroup
}

func NewTaskPool[T any](workerCount, capacity int, taskFunc func(T)) TaskPool[T] {
	ret := &taskPool[T]{
		taskQueue:   make(chan T, capacity),
		taskFunc:    taskFunc,
		workerCount: workerCount,
	}
	ret.wg.Add(workerCount)
	return ret
}

func (p *taskPool[T]) Start() {
	for i := 0; i < p.workerCount; i++ {
		go func() {
			defer p.wg.Done()

			for {
				task, taskExist := <-p.taskQueue
				if !taskExist {
					return
				}
				p.taskFunc(task)
			}
		}()
	}
}

func (p *taskPool[T]) Submit(task T) {
	p.taskQueue <- task
}

func (p *taskPool[T]) Close() {
	close(p.taskQueue)
	p.wg.Wait()
}
