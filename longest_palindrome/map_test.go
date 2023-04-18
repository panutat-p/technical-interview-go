package longest_palindrome

import "testing"

func TestLongest_empty(t *testing.T) {
	want := 0
	got := Longest("")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest_4_chars(t *testing.T) {
	want := 4
	got := Longest("aabb")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest_9_chars(t *testing.T) {
	want := 7
	got := Longest("abccccdd")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
