package transaction

import (
	"context"
	"fmt"
)

// transaction adapts a body into an executable transaction
type transaction struct {
	body   Body          // transaction's body
	done   chan struct{} // done channel is closed each time the transaction ends
	result interface{}   // transaction's result
	err    error         // if any error does happens while transaction execution it's here saved
}

func (tx *transaction) hasFinished() {
	defer close(tx.done)

	if p := recover(); p != nil {
		tx.err = fmt.Errorf(ErrTransactionPanic, p)
		tx.body.Rollback()
	} else if tx.err != nil {
		tx.body.Rollback() // err is non-nil; don't change it
	} else {
		tx.err = tx.body.Commit() // if Commit returns error update err with commit err
	}
}

func (tx *transaction) Execute(ctx context.Context) {
	defer tx.hasFinished()
	tx.done = make(chan struct{})

	if tx.err = tx.body.Precondition(); tx.err != nil {
		return
	}

	tx.result, tx.err = tx.body.Postcondition(ctx)
}

func (tx *transaction) Result() (interface{}, error) {
	return tx.result, tx.err
}

func (tx *transaction) Done() <-chan struct{} {
	return tx.done
}
