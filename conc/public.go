package conc

import (
	"sort"
	"sync"

	algr "github.com/alvidir/util/algorithm"
)

// Congruent returns a new congruent locker
func Congruent(lockers ...sync.Locker) (lockr sync.Locker) {
	if lockers == nil || len(lockers) == 0 {
		return
	}

	sort.Slice(lockers, func(i int, j int) bool {
		return algr.Address(lockers[i]) < algr.Address(lockers[i])
	})

	return &congruent{lockers: lockers}
}

// Switch switches an action to gorutine if cond is true; otherwise
// keeps in the same fiber
func Switch(cond bool, action func()) {
	if cond {
		go action()
	} else {
		action()
	}
}

// ClearMap erase all content in the sync.Map m
func ClearMap(m *sync.Map) {
	m.Range(func(key interface{}, value interface{}) bool {
		m.Delete(key)
		return true
	})
}
