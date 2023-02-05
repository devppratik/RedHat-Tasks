package strman

import (
	strman "strman/pkg"
	"testing"
)

func TestConcat(t *testing.T) {
	input := []string{"test1", "test2"}
	got := strman.ConcatStrings(input, "+")
	want := "test1+test2"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
