package parse

// Encoder represents a bilateral marshaler
type Encoder interface {
	Unmarshal() Unmarshal
	Marshal() Marshal
}
