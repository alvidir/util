package conc

import (
	"os"
	"sort"
	"sync"

	method "github.com/alvidir/util/method"
)

// Fork creates a new goroutine where to execute a provided action and returns the pid of the
// new process.
func Fork(action func()) int {
	in := make(chan int)
	go func(out chan<- int) {
		out <- os.Getpid()
		close(out)
		action()
	}(in)

	return <-in
}

// CoherentLock ensures the same locking order for a set of lockers
func CoherentLock(lockers ...sync.Locker) {
	sort.Slice(lockers[:], func(i, j int) bool {
		i_ptr, _ := method.ToUintptr(lockers[i])
		j_ptr, _ := method.ToUintptr(lockers[j])
		return i_ptr < j_ptr
	})

	for _, locker := range lockers[:] {
		if locker == nil {
			continue
		}

		locker.Lock()
	}
}

// CoherentUnlock ensures the same unlocking order for a set of lockers
func CoherentUnlock(lockers ...sync.Locker) {
	sort.Slice(lockers[:], func(i, j int) bool {
		i_ptr, _ := method.ToUintptr(lockers[i])
		j_ptr, _ := method.ToUintptr(lockers[j])
		return i_ptr < j_ptr
	})

	for _, locker := range lockers[:] {
		if locker == nil {
			continue
		}

		locker.Unlock()
	}
}

// Merge converts a list of channels to a single channel.
// For more information about Merge function go to https://blog.golang.org/pipelines.
func Merge(done <-chan struct{}, cs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	out := make(chan interface{})

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c or done is closed, then calls
	// wg.Done.
	redirect := func(in <-chan interface{}) {
		defer wg.Done()
		for got := range in {
			select {
			case out <- got:
			case <-done:
				return
			}
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go redirect(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
