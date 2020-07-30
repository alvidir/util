package sequence

import (
	"sync"
)

type Sequence struct {
	latest int64
	mu     sync.RWMutex
}

// Next returns the next int of the sequence to use
func (seq *Sequence) Next() (int64, bool) {
	seq.mu.Lock()
	defer seq.mu.Unlock()

	var ok bool
	if ok = seq.latest < seq.latest+1; ok {
		seq.latest++
	}

	return seq.latest, ok
}

// Overflow return true if, and only if, the sequence got overflow; otherwise returns false
func (seq *Sequence) Overflow() bool {
	seq.mu.RLock()
	defer seq.mu.RUnlock()

	return seq.latest+1 < seq.latest
}

// Reset sets the next value of the sequence as 0
func (seq *Sequence) Reset() {
	seq.mu.RLock()
	defer seq.mu.RUnlock()

	seq.latest = 0
}
