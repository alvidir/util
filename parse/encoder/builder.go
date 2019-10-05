package encoder

type ctrlEncoder struct {
	unmarshal Unmarshal
	marshal   Marshal
}

func (ctrl *ctrlEncoder) Unmarshal() Unmarshal {
	return ctrl.unmarshal
}

func (ctrl *ctrlEncoder) Marshal() Marshal {
	return ctrl.marshal
}

// Build builds a new Encoder for provided methods
func Build(unmarshal Unmarshal, marshal Marshal) Encoder {
	return &ctrlEncoder{
		unmarshal: unmarshal,
		marshal:   marshal,
	}
}
