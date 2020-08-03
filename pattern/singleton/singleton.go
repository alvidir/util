package singleton

import (
	"fmt"
	"sync"
)

// NewFunc is the initializer of the singleton's instance. This method is called once.
type NewFunc func() (interface{}, error)

type singleton struct {
	New      NewFunc     // singleton object creator
	instance interface{} // singleton's instance
	mu       sync.Mutex  // ensures the singleton is initialized once
}

func (s *singleton) initInstance() (err error) {
	// just when the current singleton instance is not set, the mutex has to be locked to ensure no overwrites
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.instance != nil {
		// if while waiting for mutex unlocking the instance has been initialized, then the current initialization
		// has to be suspended: singleton is ready
		return
	}

	if s.New == nil {
		// if no NewFunc has been provided for the singleton, this has no way to know how the initialization takes place.
		err = fmt.Errorf(errorNewFuncNotSet)
	}

	// otherwise create a new instance
	s.instance, err = s.New()
	return
}

func (s *singleton) GetInstance() (i interface{}, err error) {
	if i = s.instance; i != nil {
		// if singleton's instance has already been initialized: there is no sense in to locking the mutex,
		// due once the instance is successfully set, it will no longer change.
		return
	}

	if err = s.initInstance(); err != nil {
		return
	}

	i = s.instance
	return
}
