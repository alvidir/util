package counter

import "sync"

//InitCounter returns a Sequence starting from start, non default 0
func InitCounter(start int64) *Counter {
	return &Counter{latest: start}
}

// NewLockSequence returns a locking sequence
func NewLockSequence(mu sync.Locker) *Sequence {
	return &Sequence{mu: mu}
}

//InitSequence returns a Sequence starting from start, non default 0
func InitSequence(start int64) *Sequence {
	return &Sequence{latest: start}
}
