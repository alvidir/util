package util

import (
	"sync"
)

// A sequence represents a succesive
type Sequence struct {
	latest int64
	mu     sync.Mutex
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

// Reset sets the next value of the sequence as 0
func (seq *Sequence) Reset() {
	seq.mu.Lock()
	defer seq.mu.Unlock()

	seq.latest = 0
}
