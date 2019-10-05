package json

import (
	"encoding/json"

	parse "github.com/alvidir/util/parse"
)

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal parse.Unmarshal = json.Unmarshal
	return marshal.UnmarshalStrategy(filepath, manifest)
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal parse.Marshal = json.Marshal
	return marshal.MarshalStrategy(filepath, content)
}
