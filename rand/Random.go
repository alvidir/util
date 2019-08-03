package rand

import (
	"math/rand"

	operator "github.com/alvidir/util/operator"
	time "github.com/alvidir/util/time"
)

func random() *rand.Rand {
	seed := time.Unix()
	return New(seed)
}

// New builds a brand new randomizer for a given seed
func New(seed int64) *rand.Rand {
	source := rand.NewSource(seed)
	return rand.New(source)
}

// Randomize gives a pseudo-ramdom unsigned value 64 bits lenght
func Randomize() uint64 {
	return random().Uint64()
}

// Entropy returns true or false pseudo-ramdonly under the
// given entropy
func Entropy(frac float64) bool {
	limit := operator.Normalize(frac)
	cursor := random().Float64()

	return cursor <= limit
}
