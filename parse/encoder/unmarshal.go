package encoder

import "github.com/alvidir/util/stream/reader"

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
