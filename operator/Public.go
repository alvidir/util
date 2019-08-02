package operator

import (
	"math"
)

// Switch switches an action to gorutine if cond is true; otherwise
// keeps in the same fiber
func Switch(cond bool, action func()) {
	if cond {
		go action()
	} else {
		action()
	}
}

// Normalize change the state of a float to range [0.0, 1.0)
func Normalize(in float64) float64 {
	if abs := math.Abs(in); abs > 1 {
		i64 := uint64(abs)
		f64 := float64(i64)
		return abs - f64
	}

	return in
}
