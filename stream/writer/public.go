package writer

import (
	"io"
	"io/ioutil"
	"os"
)

// WriteStream returns a io.Writer for a given path
func WriteStream(rute string, build bool, append bool) (io.Writer, error) {
	flags := os.O_WRONLY
	if build { // flag de creació del fitxer si no existeix
		flags |= os.O_CREATE
	}

	if append { // flag d'escriptura continuada
		flags |= os.O_APPEND
	} else { // flag de trucament del fitxer
		flags |= os.O_TRUNC
	}

	return os.OpenFile(rute, flags, 0644)
}

// Write writes some data in to given path
func Write(filepath string, content []byte, append bool, build bool) (err error) {
	if _, err := os.Stat(filepath); err == nil || build {
		//Si el fitxer existeix, o es pretén crear-lo, procedeix l'escriptura
		err = ioutil.WriteFile(filepath, content, 0644)
	}

	return
}
