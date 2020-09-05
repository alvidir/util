package transaction

import "context"

// Body is the trait for a body
type Body interface {
	Precondition() error
	Postcondition(context.Context) (interface{}, error)
	Commit() error
	Rollback()
}

// Tx is the trait for an executable transaction
type Tx interface {
	Execute(context.Context)
	Result() (interface{}, error)
	Done() <-chan struct{}
}
