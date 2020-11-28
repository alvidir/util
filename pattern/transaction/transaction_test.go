package transaction

import (
	"context"
	"fmt"
)

const (
	iPrecondition  = "On testing precondition"
	iPostcondition = "On testing postcondition"
	iCommit        = "On testing commit"
	iRollback      = "On testing rollback"

	errPrecondition  = "Precondition failed"
	errPostcondition = "Postcondition failed"
	errCommit        = "Commit failed"
)

type testingBody struct {
	Foo          string
	FailPrecond  bool
	FailPostcond bool
	FailCommit   bool
}

func (test *testingBody) Precondition() (err error) {
	fmt.Println(iPrecondition)
	if test.FailPrecond {
		err = fmt.Errorf(errPrecondition)
	}

	return
}

func (test *testingBody) Postcondition(context.Context) (v interface{}, err error) {
	fmt.Println(iPostcondition)
	if test.FailPostcond {
		err = fmt.Errorf(errPostcondition)
	}

	v = test.Foo
	return
}

func (test *testingBody) Commit() (err error) {
	fmt.Println(iCommit)
	if test.FailCommit {
		err = fmt.Errorf(errPostcondition)
	}

	return
}

func (test *testingBody) Rollback() {
	fmt.Println(iRollback)
}

func Example_transaction_succeed() {
	want := "Example Succeed"
	subject := &testingBody{
		Foo: want,
	}

	ctx := context.Background()
	tx := NewTransaction(subject)
	tx.Execute(ctx)

	if got, err := tx.Result(); err != nil {
		fmt.Printf("Got %v, while getting transaction result\n", err.Error())
	} else if got != want {
		fmt.Printf("Got result %v, want %v\n", got, want)
	}

	// Output:
	// On testing precondition
	// On testing postcondition
	// On testing commit
}

func Example_transaction_failed() {
	subject := &testingBody{
		Foo:          "Example Failed",
		FailPostcond: true,
	}

	ctx := context.Background()
	tx := NewTransaction(subject)
	tx.Execute(ctx)

	if _, err := tx.Result(); err == nil {
		fmt.Println("Got no error on fail-forced postcondition")
	} else if err.Error() != errPostcondition {
		fmt.Printf("Got error message %v, want %v", err.Error(), errPostcondition)
	}

	// Output:
	// On testing precondition
	// On testing postcondition
	// On testing rollback
}
