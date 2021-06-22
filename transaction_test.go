package util

import (
	"context"
	"fmt"
)

func ExampleTransaction_minimal() {
	builder := NewTransactionBuilder("testing", func(Tx) (interface{}, error) {
		fmt.Println("Executing postcondition")
		return nil, nil
	})

	ctx := context.Background()
	run := builder.Build().Execute(ctx)

	// wait for the transaction to finish
	<-run.Done()

	// Output:
	// Executing postcondition
}
