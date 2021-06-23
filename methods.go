package util

import (
	"fmt"
	"strconv"
	"strings"
)

// Recoverable runs the provided method fn recovering it from panic if any.
// If panic then the trace is returned as an error, else err is nil
func Recoverable(fn func()) (err error) {
	defer func() {
		if panic := recover(); panic != nil {
			err = fmt.Errorf("%v", panic)
		}
	}()

	fn()
	return
}

// ToUintptr returns the uintptr stored into an interface.
func ToUintptr(v interface{}) (ptr uintptr, err error) {
	strptr := fmt.Sprintf("%p", v)
	strptr = strings.Replace(strptr, "0x", "", -1)

	var mempos uint64
	if mempos, err = strconv.ParseUint(strptr, 16, 64); err != nil {
		return
	}

	ptr = uintptr(mempos)
	return
}

// Find takes an slice of any type and finds the position of a provided value
func Find(slice []interface{}, val interface{}) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// FindInt takes an slice of ints and finds the position of a provided value
func FindInt(slice []int, val int) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// FindString takes an slice of strings and finds the position of a provided value
func FindString(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
