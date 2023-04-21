package make_two_anagrams

import "testing"

func TestMake1_no_action(t *testing.T) {
	want := 0
	got := Make1("anagram", "mangaar")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestMake1_chars(t *testing.T) {
	want := 1
	got := Make1("bab", "aba")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestMake1_words(t *testing.T) {
	want := 5
	got := Make1("leetcode", "practice")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
