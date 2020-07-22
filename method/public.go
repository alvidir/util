package method

import (
	"fmt"
	"strconv"
)

// Try tries to execute todo method. If panicking: the program running is restored and panic
// returned as an error. Err it's nil otherwise.
func Try(todo func()) (err error) {
	defer func() {
		if panic := recover(); panic != nil {
			err = fmt.Errorf("%v", panic)
		}
	}()

	todo()
	return
}

// ToUintptr returns the uintptr stored into an interface.
func ToUintptr(v interface{}) (ptr uintptr, err error) {
	strptr := fmt.Sprintf("%p", v)
	strptr = strptr[2:] // substring is needed to delete the 0x prefix

	var pseudo uint64
	if pseudo, err = strconv.ParseUint(strptr, 16, 32); err != nil {
		return
	}

	ptr = uintptr(pseudo)
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
