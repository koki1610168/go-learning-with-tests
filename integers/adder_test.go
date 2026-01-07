package integers

import (
	"testing"
)


func TestAdd(t *testing.T) {
	sum := Add(2, 5)
	want := 7

	if sum != want {
		t.Errorf("got %d, want %d", sum, want)
	}
}
