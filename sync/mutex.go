package sync

import (
	"sync"
	"sync/atomic"
	"unsafe"
)

const mutexLocked = 1 << iota

type mutex struct {
	sync.Mutex
}

type Mutex interface {
	Lock()
	Unlock()
	TryLock() bool
}

func NewMutex() Mutex {
	return &mutex{}
}

func (m *mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked)
}

func (m *mutex) Lock() {
	m.Mutex.Lock()
}

func (m *mutex) Unlock() {
	m.Mutex.Unlock()
}
