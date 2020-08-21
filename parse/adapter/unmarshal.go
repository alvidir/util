package adapter

import (
	"github.com/alvidir/util/stream/reader"
)

// Unmarshal is the default unmarshaler method
type Unmarshal func([]byte, interface{}) error

// Unmarshaler represents a decoder
type Unmarshaler struct {
	Fx Unmarshal
}

// Unmarshal dencode json array into an interface object
func (adapter *Unmarshaler) Unmarshal(v []byte, i interface{}) error {
	return adapter.Fx(v, i)
}

// Path decodes an encoded file to interface object
func (adapter *Unmarshaler) Path(path string, stream interface{}) (err error) {
	var content []byte
	if content, err = reader.Read(path); err == nil {
		err = adapter.Fx(content, stream)
	}

	return
}
