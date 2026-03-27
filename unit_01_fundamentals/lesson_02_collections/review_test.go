package lesson_02_collections

import "testing"

func TestSubSliceSharesBackingArray(t *testing.T) {
	nums := []int{10, 20, 30, 40, 50}
	sub := nums[1:3]
	sub[0] = 999

	if nums[1] != 999 {
		t.Fatalf("nums[1] = %d, want 999", nums[1])
	}
}

func TestMapCommaOK(t *testing.T) {
	scores := map[string]int{
		"Alice": 95,
	}

	_, ok := scores["Charlie"]
	if ok {
		t.Fatalf("expected Charlie to be absent")
	}
}
