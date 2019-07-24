package json

import (
	reader "github.com/alvidir/util/stream/reader"
	writer "github.com/alvidir/util/stream/writer"
	"encoding/json"
)

/* Unmarshal de fitxers
 * Donat el path d'un fitxer existent, en fa la reversió yaml
 * sobre la interficie passada per parametre.
 */
func Unmarshal(filepath string, stream interface{}) (err error) {
	var content []byte
	if content, err = reader.Read(filepath); err == nil {
		//Si el fitxer s'ha pogut obrir satisfactoriament
		err = json.Unmarshal(content, stream)
	}

	return
}

/* Marshal de fitxers
 * Donat el path d'un fitxer existent o no, en fa la conversió json
 * de la interficie passada per parametre i l'enmagatzema al fitxer.
 * Si aquest no existeix: el crea.
 */
func Marshal(filepath string, content *interface{}) (err error) {
	var data []byte
	if data, err = json.Marshal(content); err == nil {
		//Si la conversió ha sigut satisfactoria
		err = writer.Write(filepath, data, false, true)
	}

	return
}
