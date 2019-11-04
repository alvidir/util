package json

import (
	"encoding/json"

	adapt "github.com/alvidir/util/parse/adapter"
)

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal adapt.Unmarshal = json.Unmarshal
	return marshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal adapt.Marshal = json.Marshal
	return marshal.Path(filepath, content)
}

// Adapter builds a new adapter for json marshaling
func Adapter() adapt.Adapter {
	return adapt.New(json.Marshal, json.Unmarshal)
}
