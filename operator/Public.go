package operator

import (
	"math"
	"reflect"
)

// Abs returns the absolut value for a given int
func Abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// Clone clones an object
func Clone(in interface{}) interface{} {
	return reflect.Indirect(reflect.ValueOf(in)).Interface()
}

// Switch switches an action to gorutine if cond is true; otherwise
// keeps in the same fiber
func Switch(cond bool, action func()) {
	if cond {
		go action()
	} else {
		action()
	}
}

// Normalize change the state of a float to range [0.0, 1.0]
func Normalize(f64 float64) float64 {
	if f64 = math.Abs(f64); f64 > 1. {
		_, f64 = math.Modf(f64)
	}

	return f64
}
