package operator

type procedure func(...interface{}) error

// Ternary operates as an ternary conditional operator between
// two actions.
func Ternary(cond bool, doif procedure, ifnot procedure) error {
	if cond {
		return doif()
	}

	return ifnot()
}

// Switch switches action to gorutine if cond is true; otherwise
// keeps in the same fiber
func Switch(cond bool, action procedure) {
	if cond {
		go action()
	} else {
		action()
	}
}
