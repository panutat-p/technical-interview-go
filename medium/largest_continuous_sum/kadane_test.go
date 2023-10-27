package largest_continuous_sum

import "testing"

func TestLargestSumKadane_drop_left(t *testing.T) {
	want := 15
	got := LargestSumKadane([]int{-50, 3, 6, 4, 2})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLargestSumKadane_drop_right(t *testing.T) {
	want := 15
	got := LargestSumKadane([]int{3, 6, 4, 2, -50})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLargestSumKadane_cross(t *testing.T) {
	want := 25
	got := LargestSumKadane([]int{1, 2, -1, 3, 10, 10, -10 - 1})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
