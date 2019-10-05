package parse

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

func Build(unmarshal Unmarshal, marshal Marshal) Encoder {
	return &ctrlEncoder{
		unmarshal: unmarshal,
		marshal:   marshal,
	}
}
