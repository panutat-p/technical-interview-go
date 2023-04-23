package longest_palindrome

import "testing"

func TestLongest2_simple(t *testing.T) {
	want := 4
	got := Longest2("aabb")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest2_odd(t *testing.T) {
	want := 7
	got := Longest2("aaaaaaa")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest2_one_pair_one_char(t *testing.T) {
	want := 3
	got := Longest2("aabcdef")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest2_9_chars(t *testing.T) {
	want := 7
	got := Longest2("abccccdd")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest2_uppercase(t *testing.T) {
	want := 7
	got := Longest2("aaaAaaaa")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestLongest2_empty(t *testing.T) {
	want := 0
	got := Longest2("")
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
