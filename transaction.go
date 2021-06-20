package util

import (
	"context"
)

type Params map[string]interface{}

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
	Execute(context.Context)

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

// Transaction methods
// Execute executes the transaction for a given context
func (tx *transaction) Execute(context.Context) {

}

// Details returns the name and information of the transaction
func (tx *transaction) Details() (string, string) {
	return tx.Name, tx.Info
}

// Builder methods
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
