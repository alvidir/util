package method

import "fmt"

// Try tries to execute todo method. If panicking: the program running is restored and panic
// returned as an error. Err it's nil otherwise.
func Try(todo func()) (err error) {
	defer func() {
		if panic := recover(); panic != nil {
			err = fmt.Errorf("%v", panic)
		}
	}()

	return
}
