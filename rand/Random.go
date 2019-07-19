package rand

import (
	"math/rand"
	"time"
)

// Random genera un valor aleatori sense signe de 64 bits
func Randomize() uint64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Uint64()
}
