package strman

import (
	strman "strman/pkg"
	"testing"
)

func TestReverse(t *testing.T) {
	got := strman.Reverse("HelloThisIsATestLine")
	want := "eniLtseTAsIsihTolleH"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
