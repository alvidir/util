package operator

// Ternary operates as an ternary conditional operator between
// two actions.
func Ternary(cond bool, runmeiftrue func(), runmeifalse func()) {
	if cond {
		runmeiftrue()
	} else {
		runmeifalse()
	}
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
