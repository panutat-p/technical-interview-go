package array

import "testing"

func TestLargestSum(t *testing.T) {
	r := LargestSum([]int{1, 2, -1, 3, 4, 10, 10, -10, -1})
	if r != 29 {
		t.Error("want 29 but got", r)
	}
}

func TestLargestSum2(t *testing.T) {
	want := 25
	got := LargestSum([]int{1, 2, -1, 3, 10, 10, -10 - 1})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
