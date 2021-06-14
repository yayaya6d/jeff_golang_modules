package stack

import "sync"

type Elem interface{}

type Stack interface {
	Push(e Elem)
	Pop()
	Top() Elem
	Size() int
	IsEmpty() bool
}

type stack struct {
	elems []Elem
	mtx   sync.RWMutex
}

func NewStack() Stack {
	s := &stack{}
	s.elems = []Elem{}
	return s
}

func (s *stack) Push(e Elem) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	s.elems = append(s.elems, e)
}

func (s *stack) Pop() {
	if !s.IsEmpty() {
		s.mtx.Lock()
		defer s.mtx.Unlock()
		s.elems = s.elems[0 : len(s.elems)-1]
	}
}

func (s *stack) Top() Elem {
	if !s.IsEmpty() {
		return s.elems[len(s.elems)-1]
	}
	return nil
}

func (s *stack) Size() int {
	return len(s.elems)
}

func (s *stack) IsEmpty() bool {
	return s.Size() == 0
}
