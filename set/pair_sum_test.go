package set

import (
	"testing"
)

func TestIsPairSum(t *testing.T) {
	r := IsPairSum([]int{1, 3, 2, 2}, 4)
	if r != 2 {
		t.Error("want 2 but got", r)
	}
}
