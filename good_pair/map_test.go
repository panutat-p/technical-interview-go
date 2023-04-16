package good_pair

import "testing"

func TestCount1_none(t *testing.T) {
	want := 0
	got := Count1([]int{1, 2, 3})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestCount1_four(t *testing.T) {
	want := 4
	got := Count1([]int{1, 2, 3, 1, 1, 3})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestCount1_duplicated_four_times(t *testing.T) {
	want := 6
	got := Count1([]int{1, 1, 1, 1})
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
