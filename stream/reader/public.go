package reader

import (
	"io"
	"io/ioutil"
	"os"
)

// ReadStream returns a io.Reader for an specific path
func ReadStream(rute string) (io.Reader, error) {
	flags := os.O_RDONLY
	return os.OpenFile(rute, flags, 0644)
}

// Read returns the current content of a file
func Read(filepath string) (content []byte, err error) {
	if _, err = os.Stat(filepath); err == nil {
		//Si el fitxer existeix, procedeix la lectura
		content, err = ioutil.ReadFile(filepath)
	}

	return
}
