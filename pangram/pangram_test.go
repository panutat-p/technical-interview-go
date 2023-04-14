package pangram

import "testing"

func TestIsPangram_alphabet(t *testing.T) {
	got := IsPangram("abcdefghijklmnopqrstuvwxyz")
	if !got {
		t.Error("want true")
	}
}

func TestIsPangram_uppercase(t *testing.T) {
	got := IsPangram("TheQuickBrownFoxJumpsOverTheLazyDog")
	if !got {
		t.Error("want true")
	}
}

func TestIsPangram_space(t *testing.T) {
	got := IsPangram("This is not a pangram")
	if got {
		t.Error("want false")
	}
}

func TestIsPangram_special_char(t *testing.T) {
	got := IsPangram("example@hotmail.com")
	if got {
		t.Error("want false")
	}
}

func TestIsPangram_missing_alphabet(t *testing.T) {
	got := IsPangram("abcdefghijklmnopqrstuvw!@#")
	if got {
		t.Error("want false")
	}
}
