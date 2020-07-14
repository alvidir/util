package adapter

import "github.com/alvidir/util/stream/reader"

type Unmarshal func([]byte, interface{}) error

// UnmarshalAdapter represents a decoder
type UnmarshalAdapter struct {
	Fx Unmarshal
}

// Marshal encode an interface object into an array of bytes
func (adapter *UnmarshalAdapter) Unmarshal(v []byte, i interface{}) error {
	return adapter.Fx(v, i)
}

// Path decodes an encoded file to interface object
func (adapter *UnmarshalAdapter) Path(path string, stream interface{}) (err error) {
	var content []byte
	if content, err = reader.Read(path); err == nil {
		err = adapter.Fx(content, stream)
	}

	return
}
