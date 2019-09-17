package algorithm

import (
	"fmt"
	"testing"
)

type i interface {
	i()
}

type s struct {
}

func (str s) i() {

}

func TestAddress(t *testing.T) {
	var tester s
	var element i = tester

	fmt.Printf("Element pointer is %v\n", &element)
	fmt.Printf("Element address is %v\n", Address(&element))
}
