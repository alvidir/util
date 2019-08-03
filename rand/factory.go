package rand

import "math/rand"

// New builds a brand new randomizer for a given seed
func New(seed int64) *rand.Rand {
	source := rand.NewSource(seed)
	return rand.New(source)
}
