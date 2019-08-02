package rand

import (
	"math/rand"

	operator "github.com/alvidir/util/operator"
	time "github.com/alvidir/util/time"
)

// Random builds a brand new randomizer
func Random() *rand.Rand {
	seed := time.Unix()
	source := rand.NewSource(seed)
	return rand.New(source)
}

// Randomize gives a pseudo-ramdom unsigned value 64 bits lenght
func Randomize() uint64 {
	return Random().Uint64()
}

// Probability returns true or false pseudo-ramdonly under the
// given probability
func Probability(prob float64) bool {
	norma := operator.Normalize(prob)
	return Random().Float64() <= norma
}
