package set

import "testing"

func TestSumOfSubArray(t *testing.T) {
	r := SumOfSubArray([]int{1, 2, 3, 4, 5}, 9)
	if !r {
		t.Error("expect true")
	}
}

func TestSumOfSubArray2(t *testing.T) {
	r := SumOfSubArray([]int{1, 2, 3, 4, 5}, 8)
	if r {
		t.Error("expect false")
	}
}

func TestSumOfSubArray3(t *testing.T) {
	r := SumOfSubArray([]int{11, 9, 7, 5, 3, 2}, 9)
	if !r {
		t.Error("expect true")
	}
}

func TestSumOfSubArray4(t *testing.T) {
	r := SumOfSubArray([]int{1, -1, 3, 7, 13, 50}, 20)
	if !r {
		t.Error("expect true")
	}
}
