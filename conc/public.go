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

// ClearMap erase all content in the sync.Map m
func ClearMap(m *sync.Map) {
	m.Range(func(key interface{}, value interface{}) bool {
		m.Delete(key)
		return true
	})
}
