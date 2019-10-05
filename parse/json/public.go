package json

import (
	"encoding/json"

	coder "github.com/alvidir/util/parse/encoder"
)

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal coder.Unmarshal = json.Unmarshal
	return marshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal coder.Marshal = json.Marshal
	return marshal.Path(filepath, content)
}
