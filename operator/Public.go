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
	if f64 := math.Abs(in); f64 > 1 {
		_, in = math.Modf(f64)
	}

	return in
}
