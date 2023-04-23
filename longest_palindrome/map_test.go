package longest_palindrome

import "testing"

func TestLongest1_4_chars(t *testing.T) {
	want := 4
	got := Longest1("aabb")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest1_9_chars(t *testing.T) {
	want := 7
	got := Longest1("abccccdd")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest1_empty(t *testing.T) {
	want := 0
	got := Longest1("")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
