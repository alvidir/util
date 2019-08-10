package rand

import (
	"math/rand"

	time "github.com/alvidir/util/time"
)

// New builds a brand new randomizer for a given seed
func New(seed int64) *rand.Rand {
	source := rand.NewSource(seed)
	return rand.New(source)
}

// Random build a brand new randomizer with creation time as seed
func Random() *rand.Rand {
	seed := time.Unix()
	return New(seed)
}
