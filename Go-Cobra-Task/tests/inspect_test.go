package strman

import (
	strman "strman/pkg"
	"testing"
)

func TestInspect(t *testing.T) {
	input := "TestString123"
	gotCount, gotKind := strman.Inspect(input, false)
	wantCount := len(input)
	wantKind := "char"
	if gotCount != wantCount && gotKind != wantKind {
		t.Errorf("got %q, wanted %q", gotKind, wantKind)
		t.Errorf("got %q, wanted %q", gotCount, wantCount)
	}
}

func TestInspectWithDigit(t *testing.T) {
	input := "TestString123"
	gotCount, gotKind := strman.Inspect(input, true)
	wantCount := 3
	wantKind := "digit"
	if gotKind != wantKind {
		t.Errorf("got %q, wanted %q", gotKind, wantKind)
	}
	if gotCount != wantCount {
		t.Errorf("got %q, wanted %q", gotCount, wantCount)
	}

}
