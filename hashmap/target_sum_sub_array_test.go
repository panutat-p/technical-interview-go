package hashmap

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
