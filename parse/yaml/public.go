package yaml

import (
	code "github.com/alvidir/util/parse/encoder"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes an yaml definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var marshal code.Unmarshal = yaml.Unmarshal
	return marshal.Path(filepath, manifest)
}

// Marshal encode an interface object to corresponding yaml definition
func Marshal(filepath string, content interface{}) (err error) {
	var marshal code.Marshal = yaml.Marshal
	return marshal.Path(filepath, content)
}

// Encoder builds a new encoder for yaml marshaling
func Encoder() code.Encoder {
	return code.New(yaml.Marshal, yaml.Unmarshal)
}
