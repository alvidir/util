package encoder

type encoder struct {
	marshal   Marshal
	unmarshal Unmarshal
}

func (parse *encoder) Marshal() Marshal {
	return parse.marshal
}

func (parse *encoder) Unmarshal() Unmarshal {
	return parse.unmarshal
}

// New builds a new Encoder
func New(marshal Marshal, unmarshal Unmarshal) Encoder {
	return &encoder{marshal: marshal, unmarshal: unmarshal}
}
