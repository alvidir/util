package conc

import (
	"sort"
	"sync"

	algr "github.com/alvidir/util/algorithm"
)

// Switch switches an action to gorutine if cond is true; otherwise
// keeps in the same fiber
func Switch(cond bool, action func()) {
	if cond {
		go action()
	} else {
		action()
	}
}

// CongruentLocking locks a set of lockers in a congruent order
func CongruentLocking(lockers ...sync.Locker) {
	if lockers == nil || len(lockers) == 0 {
		return
	}

	sort.Slice(lockers, func(i int, j int) bool {
		return algr.Address(lockers[i]) < algr.Address(lockers[i])
	})

	for _, locker := range lockers {
		locker.Lock()
	}
}

// ClearMap erase all content in the sync.Map m
func ClearMap(m *sync.Map) {
	m.Range(func(key interface{}, value interface{}) bool {
		m.Delete(key)
		return true
	})
}
