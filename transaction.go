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

// A transaction represents an atomic and sequential set of actions that must be
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
	SetCommit(func(Tx)) TxBuilder
	SetRollback(func(Tx)) TxBuilder
	SetFinish(func(Tx)) TxBuilder
	//WithParameter(string, types.BasicKind) TxBuilder
	WithException(string, error) TxBuilder
	Build() Transaction
}

// NewTransactionBuilder returns a brand new builder
func NewTransactionBuilder(name string, body func(Tx) (interface{}, error)) TxBuilder {
	return &transaction{Name: name, postcondition: body}
}

type job struct {
	context.Context
	cancel context.CancelFunc
	result interface{}
	err    error
}

func (job *job) Get() (interface{}, error) {
	return job.result, job.err
}

func (job *job) Exception(string) error {
	return nil
}

func (job *job) Parameter(string) (interface{}, bool) {
	return nil, false
}

func (job *job) Sandbox(string, interface{}) (interface{}, bool) {
	return nil, false
}

func (job *job) Cancel() {
	job.cancel()
}

type transaction struct {
	Name          string
	Info          string
	precondition  func(Tx) error
	prepare       func(Tx) error
	postcondition func(Tx) (interface{}, error)
	commit        func(Tx)
	rollback      func(Tx)
	finish        func(Tx)
	//parameters    map[string]types.BasicKind
	exceptions map[string]error
}

func (tx *transaction) onFinish(job *job) {
	defer func() {
		tx.finish(job)
		job.cancel()
	}()

	if panic := recover(); panic != nil {
		job.err = fmt.Errorf("%v", panic)
		tx.rollback(job)
	} else if job.err != nil {
		tx.rollback(job) // err is non-nil; don't change it
	} else {
		tx.commit(job)
	}
}

// TRANSACTION METHODS
func (tx *transaction) Execute(ctx context.Context) Promise {
	job := &job{
		Context: ctx,
		result:  "",
	}

	defer tx.onFinish(job)
	return job
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

func (tx *transaction) SetCommit(fn func(Tx)) TxBuilder {
	tx.commit = fn
	return tx
}

func (tx *transaction) SetRollback(fn func(Tx)) TxBuilder {
	tx.rollback = fn
	return tx
}

func (tx *transaction) SetFinish(fn func(Tx)) TxBuilder {
	tx.finish = fn
	return tx
}

// func (tx *transaction) WithParameter(name string, t types.BasicKind) TxBuilder {
// 	tx.parameters[name] = t
// 	return tx
// }

func (tx *transaction) WithException(key string, err error) TxBuilder {
	tx.exceptions[key] = err
	return tx
}

func (tx *transaction) Build() Transaction {
	return tx
}
