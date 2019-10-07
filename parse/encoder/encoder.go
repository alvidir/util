package encoder

type parser struct {
	marshal   Marshal
	unmarshal Unmarshal
}

func (parse *parser) Marshal() Marshal {
	return parse.marshal
}

func (parse *parser) Unmarshal() Unmarshal {
	return parse.unmarshal
}

// New builds a new Encoder
func New(encoder Marshal, decoder Unmarshal) Encoder {
	return &parser{marshal: encoder, unmarshal: decoder}
}
