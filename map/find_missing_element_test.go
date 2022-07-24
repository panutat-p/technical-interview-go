package _map

import "testing"

// contain duplicated element
func TestFindMissing(t *testing.T) {
	r := FindMissing([]int{5, 5, 7, 7}, []int{5, 7, 7})
	if r != 5 {
		t.Error("want 5 but got", r)
	}
}

func TestFindMissing2(t *testing.T) {
	r := FindMissing([]int{3, 9, 7, 2}, []int{3, 7, 2})
	if r != 9 {
		t.Error("want 9 but got", r)
	}
}

func TestFindMissingBruteForce(t *testing.T) {
	r := FindMissingBruteForce([]int{5, 5, 5, 5, 7, 7}, []int{5, 5, 5, 7, 7})
	if r != 5 {
		t.Error("want 5 but got", r)
	}
}
