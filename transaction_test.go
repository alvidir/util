package util

import (
	"context"
	"fmt"
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
		stack = append(stack, 2)
		return nil, nil

	}).WithPrecondition(func(Tx) error {
		stack = append(stack, 1)
		return nil

	}).WithCommit(func(Tx) error {
		stack = append(stack, 3)
		return nil

	})

	ctx := context.Background()
	job := builder.Build().Execute(ctx)

	// wait for the transaction to finish
	<-job.Done()

	want := 3
	if got := len(stack); got != want {
		t.Errorf("Got transaction's execution len = %v, want %v", got, want)
	}

	for want, item := range stack {
		if got := item - 1; got != want {
			t.Errorf("Got transaction execution %v, want %v", got, want)
		}
	}

}

func TestTransaction_failed(t *testing.T) {
	forceError := func(msg string, n int, index ...int) error {
		if _, exists := FindInt(index, n); exists {
			return fmt.Errorf("error from %s", msg)
		}

		return nil
	}

	cases := []int{1, 2, 3, 4} // function where to return an error
	wants := []int{1, 3, 4, 3} //how many functions must be executed on each iteration

	for index, n := range cases {
		var stack []int
		builder := NewTransactionBuilder("testing", func(Tx) (interface{}, error) {
			stack = append(stack, 2)
			return nil, forceError("postcondition", n, 2, 4)

		}).WithPrecondition(func(Tx) error {
			stack = append(stack, 1)
			return forceError("precondition", n, 1)

		}).WithCommit(func(Tx) error {
			stack = append(stack, 3)
			return forceError("commit", n, 3)

		}).WithRollback(func(Tx) error {
			stack = append(stack, 4)
			return forceError("rollback", n, 4)

		})

		ctx := context.Background()
		job := builder.Build().Execute(ctx)

		// wait for the transaction to finish
		<-job.Done()

		if got := len(stack); got != wants[index] {
			t.Errorf("Got transaction's execution len = %v, want %v", got, wants[index])
		}

		if _, err := job.Get(); n == 4 {
			want := "error from postcondition\nerror from rollback"
			if err == nil || err.Error() != want {
				t.Errorf("Got transaction's error %v, want %v", err, want)
			}
		}

	}
}
