package set

import "testing"

func TestUniqueChar(t *testing.T) {
	want := true
	got := UniqueChar("abc")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestUniqueChar2(t *testing.T) {
	want := false
	got := UniqueChar("abcdea")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
