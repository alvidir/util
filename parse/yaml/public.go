package yaml

import (
	adapt "github.com/alvidir/util/parse/adapter"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes an yaml definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	unmarshal := adapt.NewUnmarshaler(yaml.Unmarshal)
	return unmarshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding yaml definition
func Marshal(filepath string, content interface{}) (err error) {
	marshal := adapt.NewMarshaler(yaml.Marshal)
	return marshal.Path(filepath, content)
}

// Adapter builds a new adapter for yaml marshaling
func Adapter() adapt.Adapter {
	marshal := *adapt.NewMarshaler(yaml.Marshal)
	unmarshal := *adapt.NewUnmarshaler(yaml.Unmarshal)
	return adapt.NewEncoder(marshal, unmarshal)
}
