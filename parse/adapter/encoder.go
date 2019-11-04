package adapter

type adapter struct {
	marshal   Marshal
	unmarshal Unmarshal
}

func (parse *adapter) Marshal() Marshal {
	return parse.marshal
}

func (parse *adapter) Unmarshal() Unmarshal {
	return parse.unmarshal
}

// New builds a new Encoder
func New(marshal Marshal, unmarshal Unmarshal) Adapter {
	return &adapter{marshal: marshal, unmarshal: unmarshal}
}
