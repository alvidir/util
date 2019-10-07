package encoder

import "github.com/alvidir/util/stream/writer"

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
