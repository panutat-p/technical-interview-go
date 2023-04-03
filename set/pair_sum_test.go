package set

import (
	"testing"
)

// all elements are used
func TestIsPairSum(t *testing.T) {
	r := IsPairSum([]int{1, 3, 2, 2}, 4)
	if r != 2 {
		t.Error("want 2 but got", r)
	}
}

// contain unused element
func TestIsPairSum2(t *testing.T) {
	r := IsPairSum([]int{1, 3, 2, 2, 2}, 4)
	if r != 2 {
		t.Error("want 2 but got", r)
	}
}

func TestIsPairSum3(t *testing.T) {
	r := IsPairSum([]int{4, 6, 2, 56, 8, 5, 1, 5, 6}, 10)
	if r != 3 {
		t.Error("want 3 but got", r)
	}
}

func TestIsPairSum4(t *testing.T) {
	r := IsPairSum([]int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5, 13, 14, 11, 13, -1}, 10)
	if r != 6 {
		t.Error("want 6 but got", r)
	}
}
