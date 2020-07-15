package json

import (
	"encoding/json"

	adapt "github.com/alvidir/util/parse/adapter"
)

// Unmarshal decodes an json definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	unmarshal := adapt.NewUnmarshaler(json.Unmarshal)
	return unmarshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding json definition
func Marshal(filepath string, content interface{}) (err error) {
	marshal := adapt.NewMarshaler(json.Marshal)
	return marshal.Path(filepath, content)
}

// Adapter builds a new adapter for json marshaling
func Adapter() adapt.Adapter {
	marshal := *adapt.NewMarshaler(json.Marshal)
	unmarshal := *adapt.NewUnmarshaler(json.Unmarshal)
	return adapt.NewEncoder(marshal, unmarshal)
}
