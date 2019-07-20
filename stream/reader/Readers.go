package reader

import (
	"io"
	"io/ioutil"
	"os"
)

/* Canal de lectura
 * Proporciona un canal obert de lectura per al fitxer ubicat
 * al path indicat.
 */
func ReadStream(rute string) (io.Reader, error) {
	flags := os.O_RDONLY
	return os.OpenFile(rute, flags, 0644)
}

/* Lector de fixers
 * S'encarrega de llegir el fitxer indicat pel path i
 * en retorna l'array de bytes surgent de la lectura.
 */
func Read(filepath string) (content []byte, err error) {
	if _, err = os.Stat(filepath); err == nil {
		//Si el fitxer existeix, procedeix la lectura
		content, err = ioutil.ReadFile(filepath)
	}

	return
}
