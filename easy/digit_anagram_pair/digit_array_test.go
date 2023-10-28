package digit_anagram_pair

import "testing"

func TestCount2_simple(t *testing.T) {
	want := 4
	got := Count2([]int{25, 35, 872, 228, 53, 278, 872})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount2_no(t *testing.T) {
	want := 0
	got := Count2([]int{30, 72, 3, 227})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount2_all_one(t *testing.T) {
	want := 10
	got := Count2([]int{1, 1, 1, 1, 1})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount2_zero(t *testing.T) {
	want := 1
	got := Count2([]int{0, 5, 50, 7, 0, 550, 500, 55, 11})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
