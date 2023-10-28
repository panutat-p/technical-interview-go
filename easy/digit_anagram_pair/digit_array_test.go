package digit_anagram_pair

import "testing"

func TestDigitArray_simple(t *testing.T) {
	want := 4
	got := DigitArray([]int{25, 35, 872, 228, 53, 278, 872})
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}

func TestDigitArray_no(t *testing.T) {
	want := 0
	got := DigitArray([]int{30, 72, 3, 227})
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}

func TestDigitArray_all_one(t *testing.T) {
	want := 10
	got := DigitArray([]int{1, 1, 1, 1, 1})
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}

func TestDigitArray_zero(t *testing.T) {
	want := 1
	got := DigitArray([]int{0, 5, 50, 7, 0, 550, 500, 55, 11})
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}
