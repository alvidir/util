package yaml

import (
	reader "github.com/alvidir/util/stream/reader"
	writer "github.com/alvidir/util/stream/writer"
	"gopkg.in/yaml.v2"
)

// Unmarshal decodes an yaml definition to interface object
func Unmarshal(filepath string, manifest interface{}) (err error) {
	var content []byte
	if content, err = reader.Read(filepath); err == nil {
		//Si el fitxer s'ha pogut obrir satisfactoriament
		err = yaml.Unmarshal(content, manifest)
	}

	return
}

// Marshal encode an interface object to corresponding yaml definition
func Marshal(filepath string, content interface{}) (err error) {
	var data []byte
	if data, err = yaml.Marshal(content); err == nil {
		//Si la conversió ha sigut satisfactoria
		err = writer.Write(filepath, data, false, true)
	}

	return
}
