package yaml

import (
	parse "github.com/alvidir/util/parse"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes an yaml definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal parse.Unmarshal = yaml.Unmarshal
	return marshal.UnmarshalStrategy(filepath, manifest)
}

// Marshal encode an interface object to corresponding yaml definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal parse.Marshal = yaml.Marshal
	return marshal.MarshalStrategy(filepath, content)
}
