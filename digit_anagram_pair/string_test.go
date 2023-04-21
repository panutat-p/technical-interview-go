package digit_anagram_pair

import "testing"

func TestCount1_simple(t *testing.T) {
	want := 4
	got := Count1([]int{25, 35, 872, 228, 53, 278, 872})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount1_no(t *testing.T) {
	want := 0
	got := Count1([]int{30, 72, 3, 227})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount1_all_one(t *testing.T) {
	want := 10
	got := Count1([]int{1, 1, 1, 1, 1})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount1_zero(t *testing.T) {
	want := 1
	got := Count1([]int{0, 5, 50, 7, 0, 550, 500, 55, 11})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount1_high_numbers(t *testing.T) {
	want := 2
	got := Count1([]int{327559384, 922465168, 741150091, 221559429, 903194881, 401907115, 519492225})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCount1_very_high_numbers(t *testing.T) {
	want := 2
	got := Count1([]int{692933160, 375713763, 397782110, 157470990, 235696427, 419545899, 246659723, 585206002, 731596935, 150030449, 569460123, 801563102, 287922684, 390222286, 820505620, 41869933, 666424220, 568085358, 203449619, 20313343})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
