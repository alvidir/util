package hash

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// FormatArray256 gives an string that represents the sum256 of an array of bytes
func FormatArray256(v ...byte) string {
	format := sha256.Sum256(v)
	return fmt.Sprintf("%X", format)
}

// FormatStr256 gives an string that represents the sum256 of all arguments in base 16
func FormatStr256(separator string, v ...string) string {
	join := strings.Join(v, separator)
	format := sha256.Sum256([]byte(join))

	return fmt.Sprintf("%X", format)
}
