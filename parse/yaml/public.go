package yaml

import (
	coder "github.com/alvidir/util/parse/encoder"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes an yaml definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal coder.Unmarshal = yaml.Unmarshal
	return marshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding yaml definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal coder.Marshal = yaml.Marshal
	return marshal.Path(filepath, content)
}
