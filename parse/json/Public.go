package json

import (
	"encoding/json"

	parse "github.com/alvidir/util/parse"
)

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, stream interface{}) (err error) {
	return parse.Unmarshal(filepath, stream, json.Unmarshal)
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	return parse.Marshal(filepath, content, json.Marshal)
}
