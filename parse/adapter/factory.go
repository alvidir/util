package adapter

// NewMarshaler returns the marshaler's function adapted
func NewMarshaler(marshal Marshal) *Marshaler {
	return &Marshaler{marshal}
}

// NewUnmarshaler returns the unmarshaler's function adapted
func NewUnmarshaler(unmarshal Unmarshal) *Unmarshaler {
	return &Unmarshaler{unmarshal}
}

// NewEncoder builds a new Encoder
func NewEncoder(marshal Marshaler, unmarshal Unmarshaler) Adapter {
	return &Adaptee{marshal, unmarshal}
}
