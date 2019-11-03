package conc

import (
	"os"
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
