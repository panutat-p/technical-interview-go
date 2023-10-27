package unique_char

import "testing"

func TestUniqueChar(t *testing.T) {
	want := true
	got := UniqueChar("abc")
	if got != true {
		t.Error("want", want, "but got", got)
	}
}

func TestUniqueChar2(t *testing.T) {
	want := false
	got := UniqueChar("abcdea")
	if got != false {
		t.Error("want", want, "but got", got)
	}
}
