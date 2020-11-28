package transaction

import (
	"context"
	"fmt"
)

// transaction adapts a body into an executable transaction
type transaction struct {
	body    Body          // transaction's body
	checked bool          // determines if the precondition has passed of not
	done    chan struct{} // a brand new channel is created foreach execution and closed once it ends
	result  interface{}   // transaction's result
	err     error         // if any error does happens while transaction execution it's here saved
}

func (tx *transaction) hasFinished() {
	defer func() {
		close(tx.done)
		tx.done = make(chan struct{})
	}()

	if !tx.checked {
		return
	} else if p := recover(); p != nil {
		tx.err = fmt.Errorf(ErrTransactionPanic, p)
		tx.body.Rollback()
	} else if tx.err != nil {
		tx.body.Rollback() // err is non-nil; don't change it
	} else {
		tx.err = tx.body.Commit()
	}
}

func (tx *transaction) checkBody() bool {
	if tx.err = tx.body.Precondition(); tx.err != nil {
		return false
	}

	return true
}

func (tx *transaction) Execute(ctx context.Context) {
	defer tx.hasFinished()

	if tx.checked = tx.checkBody(); !tx.checked {
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
