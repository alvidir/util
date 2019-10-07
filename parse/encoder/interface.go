package encoder

// An Encoder represents a set of de/encoder for an specific format
type Encoder interface {
	Marshal() Marshal
	Unmarshal() Unmarshal
}
