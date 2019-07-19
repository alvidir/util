package rand

import (
	"math/rand"
	time "gitlab.com/alvidir/util/time"
)

// Randomize genera un valor aleatori sense signe de 64 bits
func Randomize() uint64 {
	rand.Seed(time.Unix())
	return rand.Uint64()
}
