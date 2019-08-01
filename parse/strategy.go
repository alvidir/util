package parse

import (
	"github.com/alvidir/util/stream/reader"
	"github.com/alvidir/util/stream/writer"
)

type unmarshal func([]byte, interface{}) error
type marshal func(interface{}) ([]byte, error)

// An Encoder represents an object able to be parsed
type Encoder interface {
	Unmarshal([]byte, interface{}) error
	Marshal(interface{}) ([]byte, error)
}

// UnmarshalStrategy decodes an encoded definition to interface object
func UnmarshalStrategy(path string, stream interface{}, parser unmarshal) (err error) {
	var content []byte
	if content, err = reader.Read(path); err == nil {
		err = parser(content, stream)
	}

	return
}

// MarshalStrategy encode an interface object to corresponding coding definition
func MarshalStrategy(path string, content interface{}, parser marshal) (err error) {
	var data []byte
	if data, err = parser(content); err == nil {
		err = writer.Write(path, data, false, true)
	}

	return
}
