package sequence

import (
	"sync/atomic"
)

type Counter struct {
	count int64
}

// Get returns the current counter value
func (c *Counter) Get() int64 {
	return c.count
}

// Add increments by n the counter
func (c *Counter) Add(n int64) {
	atomic.AddInt64(&c.count, n)
}

// Increase increments by 1 the counter
func (c *Counter) Increase() {
	atomic.AddInt64(&c.count, 1)
}

// Decrease decrements by 1 the counter
func (c *Counter) Decrease() {
	atomic.AddInt64(&c.count, -1)
}
