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
