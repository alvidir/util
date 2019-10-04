package parse

// An Encoder represents an object able to be parsed
type Encoder interface {
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}
