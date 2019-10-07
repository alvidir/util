package json

import (
	"encoding/json"

	code "github.com/alvidir/util/parse/encoder"
)

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal code.Unmarshal = json.Unmarshal
	return marshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal code.Marshal = json.Marshal
	return marshal.Path(filepath, content)
}

// Encoder builds a new encoder for json marshaling
func Encoder() code.Encoder {
	return code.New(json.Marshal, json.Unmarshal)
}
