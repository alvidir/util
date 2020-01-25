package adapter

type Adaptee struct {
	X_Marshal   Marshal
	X_Unmarshal Unmarshal
}

func (parse *Adaptee) Marshal() Marshal {
	return parse.X_Marshal
}

func (parse *Adaptee) Unmarshal() Unmarshal {
	return parse.X_Unmarshal
}

// New builds a new Encoder
func New(marshal Marshal, unmarshal Unmarshal) Adapter {
	return &Adaptee{X_Marshal: marshal, X_Unmarshal: unmarshal}
}
