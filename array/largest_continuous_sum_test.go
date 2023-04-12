package array

import "testing"

func TestLargestSum_drop_left(t *testing.T) {
	want := 15
	got := LargestSum([]int{-50, 3, 6, 4, 2})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLargestSum_drop_right(t *testing.T) {
	want := 15
	got := LargestSum([]int{3, 6, 4, 2, -50})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLargestSum_cross(t *testing.T) {
	want := 25
	got := LargestSum([]int{1, 2, -1, 3, 10, 10, -10 - 1})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
