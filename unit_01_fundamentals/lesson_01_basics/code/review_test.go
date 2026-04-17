package basics

import "testing"

func TestAdd(t *testing.T) {
	got := add(10, 20)
	if got != 30 {
		t.Fatalf("add(10, 20) = %d, want 30", got)
	}
}

func TestDivMod(t *testing.T) {
	q, r := divMod(10, 3)
	if q != 3 || r != 1 {
		t.Fatalf("divMod(10, 3) = (%d, %d), want (3, 1)", q, r)
	}
}
