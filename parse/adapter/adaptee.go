package adapter

type Adaptee struct {
	X_Marshal   Marshaler
	X_Unmarshal Unmarshaler
}

func (parse *Adaptee) Encoder() Marshaler {
	return parse.X_Marshal
}

func (parse *Adaptee) Decoder() Unmarshaler {
	return parse.X_Unmarshal
}
