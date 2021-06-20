package util

import (
	"fmt"
	"sync"
)

const (
	errorNoNewFunc = "no NewFunc has been set for the singleton instance"
)

// A Singleton represents an instance of an object that has to be the same one on each call
type Singleton interface {
	// GetInstance returns the instance stored in the singleton. Multiple calls to this
	// method returns the same instance.
	GetInstance() (interface{}, error)

	// Reset forces the singleton's instance to become nil.
	Reset()
}

// NewFunc is the initializer of the singleton's instance. This method is called once.
type NewFunc func() (interface{}, error)

// A Single represents a kind of object that has to be initialized once and keep constant from there
type Single struct {
	New      NewFunc     // singleton object creator
	instance interface{} // singleton's instance
	mu       sync.Mutex  // ensures the singleton is initialized once
}

func (s *Single) initInstance() (err error) {
	// mutex locking ensure no overwrites for multiples goroutines waiting on initInstance
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.instance != nil {
		// if while waiting for mutex unlocking the instance has been initialized, then the current initialization
		// has to be suspended: singleton is ready
		return
	}

	if s.New == nil {
		// if no NewFunc has been provided for the singleton, it has no way to know how the initialization takes place.
		err = fmt.Errorf(errorNoNewFunc)
	}

	// otherwise create a new instance
	s.instance, err = s.New()
	return
}

// GetInstance returns always the same instance.
func (s *Single) GetInstance() (i interface{}, err error) {
	if i = s.instance; i != nil {
		// if singleton's instance has already been initialized: there is no sense in to locking the mutex,
		// due once the instance is successfully set, it will no longer change.
		return
	}

	// just when the current singleton instance is not set: the locking method initInstance can be called
	if err = s.initInstance(); err != nil {
		return
	}

	i = s.instance
	return
}

// Reset doesn't wait for any goroutine to end its reading, it forces the instance to become nil.
func (s *Single) Reset() {
	s.instance = nil
}
