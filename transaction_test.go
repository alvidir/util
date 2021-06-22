package util

import (
	"context"
	"testing"
)

func TestTransaction_details(t *testing.T) {
	wantName := "testing"
	wantInfo := "this is a test"

	builder := NewTransactionBuilder(wantName, func(Tx) (interface{}, error) {
		return nil, nil
	}).WithInfo(wantInfo)

	transaction := builder.Build()
	if gotName, gotInfo := transaction.Details(); gotName != wantName || gotInfo != wantInfo {
		t.Errorf("Got transaction's details %v, %v, want %v, %v", gotName, gotInfo, wantName, wantInfo)
	}
}

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
		t.Errorf("Got transaction's result %v, %v, want %v, %v", got, err, want, nil)
	}
}

func TestTransaction_success(t *testing.T) {
	var stack []int
	builder := NewTransactionBuilder("testing", func(Tx) (interface{}, error) {
		stack = append(stack, 3)
		return nil, nil
	}).WithPrecondition(func(Tx) error {
		stack = append(stack, 1)
		return nil
	}).WithPrepare(func(Tx) error {
		stack = append(stack, 2)
		return nil
	}).WithCommit(func(Tx) error {
		stack = append(stack, 4)
		return nil
	}).WithFinish(func(Tx) error {
		stack = append(stack, 5)
		return nil
	})

	ctx := context.Background()
	job := builder.Build().Execute(ctx)

	// wait for the transaction to finish
	<-job.Done()

	want := 5
	if got := len(stack); got != want {
		t.Errorf("Got transaction's execution len = %v, want %v", got, want)
	}

	for want, item := range stack {
		if got := item - 1; got != want {
			t.Errorf("Got transaction execution %v, want %v", got, want)
		}
	}

}
