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

	return uintptr(pseudo), nil
}
