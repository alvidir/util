package json

import (
	"encoding/json"

	reader "github.com/alvidir/util/stream/reader"
	writer "github.com/alvidir/util/stream/writer"
)

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, stream interface{}) (err error) {
	var content []byte
	if content, err = reader.Read(filepath); err == nil {
		err = json.Unmarshal(content, stream)
	}
	return
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	var data []byte
	if data, err = json.Marshal(content); err == nil {
		err = writer.Write(filepath, data, false, true)
	}

	return
}
