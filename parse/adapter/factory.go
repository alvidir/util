package adapter

// NewMarshaler returns the marshaler's function adapted
func NewMarshaler(marshal Marshal) Marshaler {
	return &MarshalAdapter{marshal}
}

// NewUnmarshaler returns the unmarshaler's function adapted
func NewUnmarshaler(unmarshal Unmarshal) Unmarshaler {
	return &UnmarshalAdapter{unmarshal}
}

// New builds a new Encoder
func NewEncoder(marshal Marshaler, unmarshal Unmarshaler) Adapter {
	return &Adaptee{marshal, unmarshal}
}
