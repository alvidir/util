package encoder

// An Encoder represents a set of de/encoder for an specific format
type Encoder struct {
	encode Marshal
	decode Unmarshal
}

// Marshal provides data marshaler
func (encoder *Encoder) Marshal() Marshal {
	return encoder.encode
}

// Unmarshal provides data unmarshaler
func (encoder *Encoder) Unmarshal() Unmarshal {
	return encoder.decode
}

// New builds a new Encoder
func New(encoder Marshal, decoder Unmarshal) *Encoder {
	return &Encoder{encode: encoder, decode: decoder}
}
