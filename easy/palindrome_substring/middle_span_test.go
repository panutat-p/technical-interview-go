package palindrome_substring

import (
	"testing"
)

func TestLongest1_even(t *testing.T) {
	want1 := "bab"
	want2 := "aba"
	got := Longest1("babad")
	if got != want1 && got != want2 {
		t.Error("want:", want1, "or", want2, "but got:", got)
	}
}

func TestLongest1_odd(t *testing.T) {
	want := "bb"
	got := Longest1("cbbd")
	if got != want {
		t.Error("want:", want, "but got:", got)
	}
}

func TestLongest1_even_in_odd(t *testing.T) {
	want := "bb"
	got := Longest1("abb")
	if got != want {
		t.Error("want:", want, "but got:", got)
	}
}
