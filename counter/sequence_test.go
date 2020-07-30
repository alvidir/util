package counter

import (
	"testing"
)

func TestNext(t *testing.T) {
	subject := &Sequence{}
	var iterations int64 = 100

	for want := int64(1); want < iterations; want++ {
		if got, ok := subject.Next(); got != want || !ok {
			t.Errorf("Got on next %v, %v, want %v, %v", got, ok, want, true)
		}
	}

	maxUint := ^uint64(0)
	maxInt := int64(maxUint >> 1)
	subject.latest = maxInt

	want := false
	if _, got := subject.Next(); got != want {
		t.Errorf("Got on next.ok %v, want %v", got, want)
	}
}

func TestOverflow(t *testing.T) {
	maxUint := ^uint64(0)
	maxInt := int64(maxUint >> 1)
	counter := Counter{latest: maxInt - 1}
	subject := &Sequence{Counter: counter}

	want := true
	if _, got := subject.Next(); got != want {
		t.Errorf("Got on next %v, want %v", got, want)
	}

	if got := subject.Overflow(); got != want {
		t.Errorf("Got on overflow %v, want %v", got, want)
	}
}

func TestReset(t *testing.T) {
	maxUint := ^uint64(0)
	maxInt := int64(maxUint >> 1)
	counter := Counter{latest: maxInt - 1}
	subject := &Sequence{Counter: counter}

	subject.Reset()

	var want int64 = 1
	if got, ok := subject.Next(); got != want || !ok {
		t.Errorf("Got after reset %v, %v, want %v, %v", got, ok, want, true)
	}
}
