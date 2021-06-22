package util

import (
	"context"
	"testing"
)

func TestTransaction_minimal(t *testing.T) {
	want := "hello world"
	executed := false

	builder := NewTransactionBuilder("testing", func(Tx) (interface{}, error) {
		executed = true
		return want, nil
	})

	ctx := context.Background()
	job := builder.Build().Execute(ctx)

	// wait for the transaction to finish
	<-job.Done()

	if !executed {
		t.Errorf("Got postcondition execution %v, want %v", executed, true)
	}

	if got, err := job.Get(); err != nil || got != want {
		t.Errorf("Got from transaction's result %v, %v, want %v, %v", got, err, want, nil)
	}
}
