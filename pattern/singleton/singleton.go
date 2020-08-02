package singleton

import "sync"

// NewFunc is the initializer of the singleton's instance. This method is called once.
type NewFunc func() interface{}

type singleton struct {
	New      NewFunc     // singleton object creator
	instance interface{} // singleton's instance
	once     sync.Once   // ensures the singleton is initialized once
}

func (s *singleton) init() {
	s.instance = s.New() // initiallize
}

func (s *singleton) GetInstance() interface{} {
	s.once.Do(s.init)
	return s.instance
}
