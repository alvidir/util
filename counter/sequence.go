package counter

import (
	"sync"
)

// A sequence represents a succesive
type Sequence struct {
	latest int64
	locker sync.Locker
}

// Next returns the next int of the sequence to use
func (seq *Sequence) Next() (int64, bool) {
	if locker := seq.locker; locker != nil {
		locker.Lock()
		defer locker.Unlock()
	}

	var ok bool
	if ok = seq.latest < seq.latest+1; ok {
		seq.latest++
	}

	return seq.latest, ok
}

// Overflow return true if, and only if, the sequence got overflow; otherwise returns false
func (seq *Sequence) Overflow() bool {
	if locker := seq.locker; locker != nil {
		locker.Lock()
		defer locker.Unlock()
	}

	return seq.latest+1 < seq.latest
}

// Reset sets the next value of the sequence as 0
func (seq *Sequence) Reset() {
	if locker := seq.locker; locker != nil {
		locker.Lock()
		defer locker.Unlock()
	}

	seq.latest = 0
}
