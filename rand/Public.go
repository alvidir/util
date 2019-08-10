package rand

import (
	operator "github.com/alvidir/util/operator"
)

// Range returns a pseudo-random value between [0,n]
func Range(n int) int {
	if n < 0 {
		// absolut value is needed
		n *= -1
	}

	return Random().Intn(n + 1)
}

// Int gives a pseudo-random value
func Int() int {
	return Random().Int()
}

// Uint gives a pseudo-random unsigned value
func Uint() uint {
	return uint(Random().Int())
}

// Int64 gives a pseudo-random value 64 bits lenght
func Int64() int64 {
	unsig := Random().Int63()
	if Entropy(0.5) { // fifty-fifty to become negative
		unsig *= -1
	}

	return unsig
}

// Uint64 gives a pseudo-random unsigned value 64 bits lenght
func Uint64() uint64 {
	return Random().Uint64()
}

// Entropy returns true or false pseudo-randomly under the
// given float as probability of true
func Entropy(frac float64) bool {
	switch limit := operator.Normalize(frac); {
	case limit == 0.:
		return false
	case limit == 1.:
		return true
	default:
		cursor := Random().Float64()
		return cursor <= limit
	}
}
