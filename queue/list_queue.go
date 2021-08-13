package queue

import (
	"container/list"
	"sync"
)

type ListQueue interface {
	Push(e Elem)
	Pop()
	Front() Elem
	Back() Elem
	Size() int
	IsEmpty() bool
}

type listQueue struct {
	q   *list.List
	mtx sync.RWMutex
}

func NewListQueue() ListQueue {
	lq := &listQueue{
		q: list.New(),
	}
	return lq
}

func (lq *listQueue) Size() int {
	return lq.q.Len()
}

func (lq *listQueue) IsEmpty() bool {
	return lq.q.Len() == 0
}

func (lq *listQueue) Push(e Elem) {
	lq.mtx.Lock()
	defer lq.mtx.Unlock()
	lq.q.PushBack(e)
}

func (lq *listQueue) Pop() {
	if !lq.IsEmpty() {
		lq.mtx.Lock()
		defer lq.mtx.Unlock()
		lq.q.Remove(lq.q.Front())
	}
}

func (lq *listQueue) Front() Elem {
	if lq.IsEmpty() {
		return nil
	}
	return lq.q.Front().Value
}

func (lq *listQueue) Back() Elem {
	if lq.IsEmpty() {
		return nil
	}
	return lq.q.Back().Value
}
