package yaml

import (
	adapt "github.com/alvidir/util/parse/adapter"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes an yaml definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal adapt.Unmarshal = yaml.Unmarshal
	return marshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding yaml definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal adapt.Marshal = yaml.Marshal
	return marshal.Path(filepath, content)
}

// Adapter builds a new adapter for yaml marshaling
func Adapter() adapt.Adapter {
	return adapt.New(yaml.Marshal, yaml.Unmarshal)
}
