package util

import (
	"context"
	"fmt"
)

type Params map[string]interface{}

// A Promise represents some value that will be resolved in the future
type Promise interface {
	// Done returns a read-only channel that will be closed as soon as the
	// value gets resolved or an error happens
	Done() <-chan struct{}

	// Get returns the expected value or and error, if any
	Get() (interface{}, error)
}

// A Tx is the context of a transaction's execution
type Tx interface {
	context.Context

	// Exception crashes the transaction's execution if stills running
	// and the exception exists, otherwise err != nil
	Exception(string) error

	// Parameter takes the value with key 'data' from the context, expecting an
	// object of type Params and looks up for the provided key
	Parameter(string) (interface{}, bool)

	// Sandbox loads and stores data under a key in order to make it reachable
	// for other steps of the same execution. If an string is provided with no
	// value/interface then the function behaves as a getter, otherwise it stores
	// the value an returns the old one, if any
	Sandbox(string, interface{}) (interface{}, bool)

	// Cancel cancels the transaction's context, stoping its exeution as well
	Cancel()
}

// A Transaction represents an atomic and sequential set of actions that must be
// fully performed or reversed otherwise
type Transaction interface {
	// Execute executes the transaction for a given context
	Execute(context.Context) Promise

	// Details returns the name and information of the transaction
	Details() (string, string)
}

type TxBuilder interface {
	SetInfo(string) TxBuilder
	SetPrecondition(func(Tx) error) TxBuilder
	SetPrepare(func(Tx) error) TxBuilder
	SetCommit(func(Tx) error) TxBuilder
	SetRollback(func(Tx) error) TxBuilder
	SetFinish(func(Tx) error) TxBuilder
	WithException(string, error) TxBuilder
	Build() Transaction
}

// NewTransactionBuilder returns a brand new builder
func NewTransactionBuilder(name string, body func(Tx) (interface{}, error)) TxBuilder {
	return &transaction{Name: name, postcondition: body}
}

type thread struct {
	context.Context
	*transaction

	cancel context.CancelFunc
	result interface{}
	err    error
}

func (thread *thread) Get() (interface{}, error) {
	return thread.result, thread.err
}

func (thread *thread) Exception(string) error {
	return nil
}

func (thread *thread) Parameter(string) (interface{}, bool) {
	return nil, false
}

func (thread *thread) Sandbox(string, interface{}) (interface{}, bool) {
	return nil, false
}

func (thread *thread) Cancel() {
	thread.cancel()
}

type transaction struct {
	Name          string
	Info          string
	precondition  func(Tx) error
	prepare       func(Tx) error
	postcondition func(Tx) (interface{}, error)
	commit        func(Tx) error
	rollback      func(Tx) error
	finish        func(Tx) error
	exceptions    map[string]error
}

func (tx *transaction) doRollback(thread *thread) {
	if err := tx.rollback(thread); err != nil {
		if thread.err != nil {
			err = fmt.Errorf("%s\n%s", thread.err.Error(), err.Error())
		}

		thread.err = err
	}
}

func (tx *transaction) doCommit(thread *thread) {
	if thread.err = tx.commit(thread); thread.err != nil {
		tx.doRollback(thread)
	}
}

func (tx *transaction) finalize(thread *thread) {
	defer tx.done(thread)

	if panic := recover(); panic != nil {
		thread.err = fmt.Errorf("%v", panic)
		tx.doRollback(thread)
	} else if thread.err != nil {
		tx.doRollback(thread) // err is non-nil; don't change it
	} else {
		tx.doCommit(thread)
	}
}

func (tx *transaction) done(thread *thread) {
	if tx.finish != nil {
		// finish is an optional function, so its nullability must be checked
		if err := tx.finish(thread); err != nil {
			if thread.err != nil {
				err = fmt.Errorf("%s\n%s", thread.err.Error(), err.Error())
			}

			thread.err = err
		}
	}

	thread.cancel()
}

func (tx *transaction) run(thread *thread) {
	if tx.precondition != nil {
		// precondition is an optional function, so its nullability must be checked
		if thread.err = tx.precondition(thread); thread.err != nil {
			thread.cancel()
			return
		}
	}

	defer tx.finalize(thread)

	if tx.prepare != nil {
		// prepare is an optional function, so its nullability must be checked
		tx.prepare(thread)
	}

	if thread.result, thread.err = tx.postcondition(thread); thread.err != nil {
		thread.cancel()
		return
	}
}

// TRANSACTION METHODS
func (tx *transaction) Execute(ctx context.Context) Promise {
	ctx, cancel := context.WithCancel(ctx)

	thread := &thread{
		Context:     ctx,
		transaction: tx,
		cancel:      cancel,
	}

	go tx.run(thread)
	return thread
}

func (tx *transaction) Details() (string, string) {
	return tx.Name, tx.Info
}

// BUILDER METHODS
func (tx *transaction) SetInfo(info string) TxBuilder {
	tx.Info = info
	return tx
}

func (tx *transaction) SetPrecondition(fn func(Tx) error) TxBuilder {
	tx.precondition = fn
	return tx
}

func (tx *transaction) SetPrepare(fn func(Tx) error) TxBuilder {
	tx.prepare = fn
	return tx
}

func (tx *transaction) SetCommit(fn func(Tx) error) TxBuilder {
	tx.commit = fn
	return tx
}

func (tx *transaction) SetRollback(fn func(Tx) error) TxBuilder {
	tx.rollback = fn
	return tx
}

func (tx *transaction) SetFinish(fn func(Tx) error) TxBuilder {
	tx.finish = fn
	return tx
}

func (tx *transaction) WithException(key string, err error) TxBuilder {
	tx.exceptions[key] = err
	return tx
}

func (tx *transaction) Build() Transaction {
	return tx
}
