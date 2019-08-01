package yaml

import (
	parse "github.com/alvidir/util/parse"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes an yaml definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	return parse.UnmarshalStrategy(filepath, manifest, yaml.Unmarshal)
}

// Marshal encode an interface object to corresponding yaml definition
func Marshal(filepath string, content interface{}) (err error) {
	return parse.Marshal(filepath, content, yaml.Marshal)
}
