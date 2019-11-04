package adapter

// An Adapter represents a set of de/encoder for an specific format
type Adapter interface {
	Marshal() Marshal
	Unmarshal() Unmarshal
}
