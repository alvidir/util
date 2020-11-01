package transaction

// NewTransaction builds an executable transaction with the given body
func NewTransaction(body Body) Tx {
	return &transaction{
		body: body,
		done: make(chan struct{}),
	}
}
