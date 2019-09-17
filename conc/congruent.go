package conc

type congruent struct {
	lockers []sync.Locker
}

func (congr *congruent) Lock() {
	for _, locker := range congr.lockers {
		locker.Lock()
	}
}

func (congr *congruent) Unlock() {
	for _, locker := range congr.lockers {
		locker.Unlock()
	}
}