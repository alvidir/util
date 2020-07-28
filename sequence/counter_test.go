package sequence

import (
	"testing"
)

func TestCounter(t *testing.T) {
	subject := &Counter{}

	want := 0
	if got := subject.Get(); got != want {
		t.Errorf("Got on count %v, want %v", got, want)
	}

	want = 100
	subject.Add(want)
	if got := subject.Get(); got != want {
		t.Errorf("Got on count %v, want %v", got, want)
	}

	want = 50
	subject.Add(-want)
	if got := subject.Get(); got != want {
		t.Errorf("Got on count %v, want %v", got, want)
	}
}