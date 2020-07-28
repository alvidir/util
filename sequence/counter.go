package sequence

import (
	"sync/atomic"
)

type Counter struct {
	count int32
}

// Get returns the current counter value
func (c *Counter) Get() int {
	return int(c.count)
}

// Add increments by n the counter
func (c *Counter) Add(n int) {
	atomic.AddInt32(&c.count, int32(n))
}

// Increase increments by 1 the counter
func (c *Counter) Increase() {
	atomic.AddInt32(&c.count, 1)
}

// Decrease decrements by 1 the counter
func (c *Counter) Decrease() {
	atomic.AddInt32(&c.count, -1)
}
