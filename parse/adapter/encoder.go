package adapter

type adapter struct {
	X_Marshal   Marshal
	X_Unmarshal Unmarshal
}

func (parse *adapter) Marshal() Marshal {
	return parse.X_Marshal
}

func (parse *adapter) Unmarshal() Unmarshal {
	return parse.X_Unmarshal
}

// New builds a new Encoder
func New(marshal Marshal, unmarshal Unmarshal) Adapter {
	return &adapter{X_Marshal: marshal, X_Unmarshal: unmarshal}
}
