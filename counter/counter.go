package counter

import (
	"sync/atomic"
)

type Counter struct {
	latest int64
}

// Get returns the current counter value
func (c *Counter) Get() int64 {
	return atomic.LoadInt64(&c.latest)
}

// Add increments by n the counter
func (c *Counter) Add(n int64) {
	atomic.AddInt64(&c.latest, n)
}

// Increase increments by 1 the counter
func (c *Counter) Increase() {
	atomic.AddInt64(&c.latest, 1)
}

// Decrease decrements by 1 the counter
func (c *Counter) Decrease() {
	atomic.AddInt64(&c.latest, -1)
}
