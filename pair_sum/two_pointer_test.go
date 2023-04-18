package pair_sum

import (
	"reflect"
	"testing"
)

func TestSum1_equal_6(t *testing.T) {
	want := []int{1, 3}
	got := Sum1([]int{1, 2, 3, 4, 6}, 6)
	if !reflect.DeepEqual(want, got) {
		t.Error("want", want, "but got", got)
	}
}

func TestSum2_equal_11(t *testing.T) {
	want := []int{0, 2}
	got := Sum1([]int{2, 5, 9, 11}, 11)
	if !reflect.DeepEqual(want, got) {
		t.Error("want", want, "but got", got)
	}
}
