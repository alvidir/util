package sequence

import (
	"sync"
	"sync/atomic"
)

type sequence struct {
	latest int
	count  int32
	mu     sync.Mutex
}

// Next returns the next int of the sequence to use
func (seq *sequence) Next() (int, bool) {
	seq.mu.Lock()
	defer seq.mu.Unlock()

	var ok bool
	if ok = seq.latest < seq.latest+1; ok {
		seq.latest++
	}

	return seq.latest, ok
}

// Overflow return true if, and only if, the sequence got overflow; otherwise returns false
func (seq *sequence) Overflow() bool {
	return seq.latest+1 < seq.latest
}

// Count returns the current counter state
func (seq *sequence) Count() int {
	return int(seq.count)
}

// Add increments by n the sequence counter
func (seq *sequence) Add(n int) {
	atomic.AddInt32(&seq.count, int32(n))
}

// Reset sets the next value of the sequence as 0
func (seq *sequence) Reset() {
	seq.latest = 0
}
