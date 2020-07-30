package counter

import "sync"

//InitCounter returns a new counter starting from start, non default 0
func InitCounter(start int64) *Counter {
	return &Counter{latest: start}
}

// NewLockSequence returns a locking sequence
func NewLockSequence(mu sync.Locker) *Sequence {
	return &Sequence{locker: mu}
}

//InitSequence returns a sequence starting from start, non default 0
func InitSequence(start int64) *Sequence {
	return &Sequence{latest: start}
}
