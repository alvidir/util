package pattern

// A Tx interface represents a non return executable instance that
// keeps self contained
type Tx interface {
	Execute()
}
