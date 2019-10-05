package encoder

import (
	"github.com/alvidir/util/stream/reader"
	"github.com/alvidir/util/stream/writer"
)

// Unmarshal represents a decoder
type Unmarshal func([]byte, interface{}) error

// Path decodes an encoded file to interface object
func (marshal Unmarshal) Path(path string, stream interface{}) (err error) {
	var content []byte
	if content, err = reader.Read(path); err == nil {
		err = marshal(content, stream)
	}

	return
}

// Marshal represents a encoder
type Marshal func(interface{}) ([]byte, error)

// Path encode an interface object to corresponding path
func (marshal Marshal) Path(path string, content interface{}) (err error) {
	var data []byte
	if data, err = marshal(content); err == nil {
		err = writer.Write(path, data, false, true)
	}

	return
}
