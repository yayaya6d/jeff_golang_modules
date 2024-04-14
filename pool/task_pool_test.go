package pool

import (
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPool_GivenSomeTask_FinishProperly(t *testing.T) {
	var count atomic.Int32
	taskFunc := func(i int32) {
		count.Add(i)
	}

	p := NewTaskPool(10, 10, taskFunc)
	p.Start()
	for i := 1; i <= 10; i++ {
		p.Submit(int32(i))
	}
	p.Close()

	assert.Equal(t, int32(55), count.Load())
}
