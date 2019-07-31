package operator

type function func(...interface{}) error
type procedure func()

// Ternary operates as an ternary conditional operator between
// two actions.
func Ternary(cond bool, doif function, ifnot function) error {
	if cond {
		return doif()
	}

	return ifnot()
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
