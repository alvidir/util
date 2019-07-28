package parse

import (
	"github.com/alvidir/util/stream/reader"
	"github.com/alvidir/util/stream/writer"
)

// An Encoder represents an object able to be parsed
type Encoder interface {
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}

// UnmarshalStrategy decodes an encoded definition to interface object
func UnmarshalStrategy(filepath string, stream interface{}, code Encoder) (err error) {
	var content []byte
	if content, err = reader.Read(filepath); err == nil {
		err = code.Unmarshal(content, stream)
	}

	return
}

// MarshalStrategy encode an interface object to corresponding coding definition
func MarshalStrategy(filepath string, content interface{}, code Encoder) (err error) {
	var data []byte
	if data, err = code.Marshal(content); err == nil {
		err = writer.Write(filepath, data, false, true)
	}

	return
}
