package util

import (
	"context"
	"fmt"
	"sync"
)

const (
	TxData              Key    = "data"
	ErrUnknownException string = "unknown exception"
)

type Key string
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

	// Exception cancels the context of the calling thread if stills running
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
	WithInfo(string) TxBuilder
	WithPrecondition(func(Tx) error) TxBuilder
	WithCommit(func(Tx) error) TxBuilder
	WithRollback(func(Tx) error) TxBuilder
	WithException(string, error) TxBuilder
	Build() Transaction
}

// NewTransactionBuilder returns a brand new transaction builder
func NewTransactionBuilder(name string, body func(Tx) (interface{}, error)) TxBuilder {
	return &transaction{Name: name, postcondition: body}
}

type thread struct {
	sync.Mutex
	context.Context
	*transaction

	cancel  context.CancelFunc
	sandbox map[string]interface{}
	result  interface{}
	err     error
}

func (thread *thread) error(err error) {
	thread.Lock()
	defer thread.Unlock()

	if thread.err != nil && err != nil {
		thread.err = fmt.Errorf("%s\n%s", thread.err.Error(), err.Error())
	} else if err != nil {
		thread.err = err
	}
}

func (thread *thread) Get() (interface{}, error) {
	thread.Lock()
	defer thread.Unlock()
	return thread.result, thread.err
}

func (thread *thread) Exception(id string) error {
	v, exists := thread.exceptions.Load(id)
	if !exists {
		return fmt.Errorf(ErrUnknownException)
	}

	thread.error(v.(error))
	thread.cancel()
	return nil
}

func (thread *thread) Parameter(key string) (v interface{}, ok bool) {
	data := thread.Value(TxData)
	if ok = data != nil; !ok {
		return
	}

	var params Params
	if params, ok = data.(Params); !ok {

		return
	}

	v, ok = params[key]
	return
}

func (thread *thread) Sandbox(key string, new interface{}) (old interface{}, ok bool) {
	if old, ok = thread.sandbox[key]; new != nil {
		// if a new value is set, then the method behaves as a Load and Store function
		thread.sandbox[key] = new
	}

	return
}

func (thread *thread) Cancel() {
	thread.cancel()
}

type transaction struct {
	Name          string
	Info          string
	precondition  func(Tx) error
	postcondition func(Tx) (interface{}, error)
	commit        func(Tx) error
	rollback      func(Tx) error
	exceptions    sync.Map
}

func (tx *transaction) doRollback(thread *thread) {
	if tx.rollback != nil { // rollback is optional
		if err := tx.rollback(thread); err != nil {
			thread.error(err)
		}
	}
}

func (tx *transaction) doCommit(thread *thread) {
	if tx.commit != nil { // commit is optional
		if thread.err = tx.commit(thread); thread.err != nil {
			tx.doRollback(thread)
		}
	}
}

func (tx *transaction) finalize(thread *thread) {
	defer thread.cancel()

	// if the context has been canceled it is required to catch the cause
	err := thread.Err()
	thread.error(err)

	if panic := recover(); panic != nil {
		thread.err = fmt.Errorf("%v", panic)
		tx.doRollback(thread)
	} else if thread.err != nil {
		tx.doRollback(thread) // err is non-nil; don't change it
	} else {
		tx.doCommit(thread)
	}
}

func (tx *transaction) spawn(thread *thread) <-chan struct{} {
	done := make(chan struct{})
	go func(done chan<- struct{}) {
		defer close(done)

		if result, err := tx.postcondition(thread); err == nil {
			thread.Lock()
			defer thread.Unlock()
			thread.result = result
		} else {
			thread.error(err)
		}

	}(done)

	return done
}

func (tx *transaction) run(thread *thread) {
	if tx.precondition != nil { // precondition is optional
		if err := tx.precondition(thread); err != nil {
			thread.error(err)
			thread.cancel()
			return
		}
	}

	defer tx.finalize(thread)

	select {
	case <-tx.spawn(thread):
		// the body of the transaction has been completed successfully
	case <-thread.Done():
		// the context has been canceled before the body could finish
		// so the execution must be stoped
	}
}

// TRANSACTION METHODS
func (tx *transaction) Execute(ctx context.Context) Promise {
	ctx, cancel := context.WithCancel(ctx)

	thread := &thread{
		Context:     ctx,
		transaction: tx,
		cancel:      cancel,
		sandbox:     make(map[string]interface{}),
	}

	go tx.run(thread)
	return thread
}

func (tx *transaction) Details() (string, string) {
	return tx.Name, tx.Info
}

// BUILDER METHODS
func (tx *transaction) WithInfo(info string) TxBuilder {
	tx.Info = info
	return tx
}

func (tx *transaction) WithPrecondition(fn func(Tx) error) TxBuilder {
	tx.precondition = fn
	return tx
}

func (tx *transaction) WithCommit(fn func(Tx) error) TxBuilder {
	tx.commit = fn
	return tx
}

func (tx *transaction) WithRollback(fn func(Tx) error) TxBuilder {
	tx.rollback = fn
	return tx
}

func (tx *transaction) WithException(key string, err error) TxBuilder {
	tx.exceptions.Store(key, err)
	return tx
}

func (tx *transaction) Build() Transaction {
	// the build function shall clone the transaction itself in order to avoi potential
	// modifications of the transaction once builded
	clone := &transaction{
		Name:          tx.Name,
		Info:          tx.Info,
		precondition:  tx.precondition,
		postcondition: tx.postcondition,
		commit:        tx.commit,
		rollback:      tx.rollback,
		exceptions:    tx.exceptions,
	}

	return clone
}
