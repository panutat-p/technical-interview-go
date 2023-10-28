package digit_anagram_pair

import "testing"

func TestSortedDigits(t *testing.T) {
	t.Run("4 pairs", func(t *testing.T) {
		want := 4
		got := SortedDigits([]int{25, 35, 872, 228, 53, 278, 872})
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("no pair", func(t *testing.T) {
		want := 0
		got := SortedDigits([]int{30, 72, 3, 227})
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("all one", func(t *testing.T) {
		want := 10
		got := SortedDigits([]int{1, 1, 1, 1, 1})
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("zero pair", func(t *testing.T) {
		want := 1
		got := SortedDigits([]int{0, 5, 50, 7, 0, 550, 500, 55, 11})
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("high numbers", func(t *testing.T) {
		want := 2
		got := SortedDigits([]int{327559384, 922465168, 741150091, 221559429, 903194881, 401907115, 519492225})
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("big numbers", func(t *testing.T) {
		want := 2
		got := SortedDigits([]int{692933160, 375713763, 397782110, 157470990, 235696427, 419545899, 246659723, 585206002, 731596935, 150030449, 569460123, 801563102, 287922684, 390222286, 820505620, 41869933, 666424220, 568085358, 203449619, 20313343})
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
}
