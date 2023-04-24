package palindrome_substring

import (
	"testing"
)

func TestLongest1_simple(t *testing.T) {
	want1 := "bab"
	want2 := "aba"
	got := Longest1("babad")
	if got != want1 && got != want2 {
		t.Error("want:", want1, "or", want2, "but got:", got)
	}
}
