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

// Int gives a pseudo-random value
func Int() int {
	return random().Int()
}

// Uint gives a pseudo-random unsigned value
func Uint() uint {
	return uint(random().Int())
}

// Uint64 gives a pseudo-random unsigned value 64 bits lenght
func Uint64() uint64 {
	return random().Uint64()
}

// Entropy returns true or false pseudo-randomly under the
// given float as probability of true
func Entropy(frac float64) bool {
	bounds := operator.Normalize(frac)
	cursor := random().Float64()

	return cursor <= bounds
}
