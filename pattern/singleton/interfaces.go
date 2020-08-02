package singleton

// A Singleton represents an instance of an object that has to be the same one on each call
type Singleton interface {
	// GetInstance returns the instance stored in the singleton. Multiple calls to this
	// method returns the same instance.
	GetInstance() interface{}
}