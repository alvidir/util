package hash

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// FormatStr256 gives an string that represents the sum256 of all arguments in base 16
func FormatStr256(args ...string) (hash string) {
	hash = strings.Join(args, "")
	argot := []byte(hash)
	forma := sha256.Sum256(argot)

	return fmt.Sprintf("%X", forma)
}
