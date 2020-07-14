package adapter

type stream interface {
	Path(string, interface{}) error
}

// An Adapter represents a set of de/encoder for an specific format
type Adapter interface {
	Encoder() Marshaler
	Decoder() Unmarshaler
}

// Marshaler represents a set of ways for encoding
type Marshaler interface {
	stream
	Marshal(interface{}) ([]byte, error)
}

// Unmarshaler represents a set of ways for dencoding
type Unmarshaler interface {
	stream
	Unmarshal([]byte, interface{}) error
}
