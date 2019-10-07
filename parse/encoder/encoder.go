package encoder

type ctrlEncoder struct {
	marshal   Marshal
	unmarshal Unmarshal
}

func (parse *ctrlEncoder) Marshal() Marshal {
	return parse.marshal
}

func (parse *ctrlEncoder) Unmarshal() Unmarshal {
	return parse.unmarshal
}

// New builds a new Encoder
func New(encoder Marshal, decoder Unmarshal) Encoder {
	return &ctrlEncoder{marshal: encoder, unmarshal: decoder}
}
