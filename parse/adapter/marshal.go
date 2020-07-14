package adapter

import "github.com/alvidir/util/stream/writer"

type Marshal func(interface{}) ([]byte, error)

// MarshalAdapter represents a encoder
type MarshalAdapter struct {
	Fx Marshal
}

// Marshal encode an interface object into an array of bytes
func (adapter *MarshalAdapter) Marshal(i interface{}) ([]byte, error) {
	return adapter.Fx(i)
}

// Path encode an interface object to corresponding path
func (adapter *MarshalAdapter) Path(path string, content interface{}) (err error) {
	var data []byte
	if data, err = adapter.Fx(content); err == nil {
		err = writer.Write(path, data, false, true)
	}

	return
}
