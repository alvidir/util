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

// CloneMap returns a copy of a provided sync.Map into a map[interface{}]interface{}
func CloneMap(in *sync.Map) (out map[interface{}]interface{}) {
	if in == nil {
		return
	}

	in.Range(func(key interface{}, value interface{}) bool {
		out[key] = value
		return true
	})

	return out
}
