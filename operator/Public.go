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
func Clone(original interface{}) interface{} {
	clone := reflect.ValueOf(original)
	clone = reflect.Indirect(clone)
	return clone.Interface()
}

// Normalize change the state of a float to range [0.0, 1.0]
func Normalize(f64 float64) float64 {
	if f64 = math.Abs(f64); f64 > 1. {
		_, f64 = math.Modf(f64)
	}

	return f64
}
