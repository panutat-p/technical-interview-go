package longest_palindrome

import "testing"

func TestLongest1_simple(t *testing.T) {
	want := 4
	got := Longest1("aabb")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest1_odd(t *testing.T) {
	want := 7
	got := Longest1("aaaaaaa")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest1_one_pair_one_char(t *testing.T) {
	want := 3
	got := Longest1("aabcdef")
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

func TestLongest1_uppercase(t *testing.T) {
	want := 7
	got := Longest1("aaaAaaaa")
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
