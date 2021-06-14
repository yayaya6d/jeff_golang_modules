package queue

import "sync"

type Elem interface{}

type Queue interface {
	Push(e Elem)
	Pop()
	Front() Elem
	Back() Elem
	Size() int
	IsEmpty() bool
}

type queue struct {
	elems []Elem
	mtx   sync.RWMutex
}

func NewQueue() Queue {
	q := &queue{
		elems: []Elem{},
	}
	return q
}

func (q *queue) Size() int {
	return len(q.elems)
}

func (q *queue) IsEmpty() bool {
	return q.Size() == 0
}

func (q *queue) Push(e Elem) {
	q.mtx.Lock()
	defer q.mtx.Unlock()
	q.elems = append(q.elems, e)
}

func (q *queue) Pop() {
	if !q.IsEmpty() {
		q.mtx.Lock()
		defer q.mtx.Unlock()
		q.elems = q.elems[1:len(q.elems)]
	}
}

func (q *queue) Front() Elem {
	if q.IsEmpty() {
		return nil
	}
	return q.elems[0]
}

func (q *queue) Back() Elem {
	if q.IsEmpty() {
		return nil
	}
	return q.elems[len(q.elems)-1]
}
