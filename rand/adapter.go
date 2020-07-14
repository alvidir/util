package rand

import (
	"math/rand"

	operator "github.com/alvidir/util/operator"
)

type Adapter struct {
	*rand.Rand
}

// Range returns a pseudo-random value between [0,n]
func (rand *Adapter) Range(n int) int {
	if n < 0 {
		// absolut value is needed
		n *= -1
	}

	return rand.Intn(n + 1)
}

// Uint gives a pseudo-random unsigned value
func (rand *Adapter) Uint() uint {
	return uint(rand.Int())
}

// Int64 gives a pseudo-random value 64 bits lenght
func (rand *Adapter) Int64() int64 {
	unsig := rand.Int63()
	if rand.Entropy(0.5) { // fifty-fifty to become negative
		unsig *= -1
	}

	return unsig
}

// Entropy returns true or false pseudo-randomly under the
// given float as probability of true
func (rand *Adapter) Entropy(frac float64) bool {
	switch limit := operator.Normalize(frac); {
	case limit == 0.:
		return false
	case limit == 1.:
		return true
	default:
		cursor := rand.Float64()
		return cursor <= limit
	}
}
