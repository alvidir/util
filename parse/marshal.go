package parse

import (
	"github.com/alvidir/util/stream/reader"
	"github.com/alvidir/util/stream/writer"
)

// Unmarshal represents a decoder
type Unmarshal func([]byte, interface{}) error

// Marshal represents a encoder
type Marshal func(interface{}) ([]byte, error)

// UnmarshalStrategy decodes an encoded definition to interface object
func (marshal Unmarshal) UnmarshalStrategy(path string, stream interface{}) (err error) {
	var content []byte
	if content, err = reader.Read(path); err == nil {
		err = marshal(content, stream)
	}

	return
}

// MarshalStrategy encode an interface object to corresponding coding definition
func (marshal Marshal) MarshalStrategy(path string, content interface{}) (err error) {
	var data []byte
	if data, err = marshal(content); err == nil {
		err = writer.Write(path, data, false, true)
	}

	return
}
