package operator

type procedure func(...interface{}) error

// Ternary operates as an ternary conditional operator for
func Ternary(cond bool, doif procedure, ifnot procedure) error {
	if cond {
		return doif()
	}

	return ifnot()
}
