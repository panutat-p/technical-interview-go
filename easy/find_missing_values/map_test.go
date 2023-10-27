package find_missing_values

import "testing"

func TestFind1_missing_five(t *testing.T) {
	want := 5
	got := Find1([]int{5, 5, 7, 7}, []int{5, 7, 7})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestFind1_missing_nine(t *testing.T) {
	want := 9
	got := Find1([]int{3, 9, 7, 2}, []int{3, 7, 2})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
