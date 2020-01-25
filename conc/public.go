package conc

import (
	"os"
	"sync"
)

// Switch switches an action to gorutine and returns the new process pid if cond is true;
// otherwise keeps in the same fiber
func Switch(cond bool, action func()) int {
	if !cond {
		action()
		return 0
	}

	pid := make(chan int)
	go func(chann chan<- int) {
		pid <- os.Getpid()
		action()
	}(pid)

	return <-pid
}

// Coherence ensures the same lock/unlock order for a set of lockers, being lock == true if,
// and only if, the function to execute for each locker is Lock(), otherwise it's Unlock().
func Coherence(lock bool, lockers ...sync.Locker) {

}

// Merge converts a list of channels to a single channel
// For more information about Merge function go to https://blog.golang.org/pipelines
func Merge(done <-chan struct{}, cs ...<-chan interface{}) <-chan interface{} {
	var wg sync.WaitGroup
	output := make(chan interface{})

	wg.Add(len(cs))
	for _, c := range cs {
		// Start an output goroutine for each input channel in cs.  output
		// copies values from c to out until c or done is closed, then calls
		// wg.Done.
		go func(in <-chan interface{}, out chan<- interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
			defer wg.Done()
			for got := range in {
				select {
				case out <- got:
				case <-done:
					return
				}
			}
		}(c, output, done, &wg)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(output)
	}()
	return output
}
