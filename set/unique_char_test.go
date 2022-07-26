package set

import "testing"

func TestUniqueChar(t *testing.T) {
	r := UniqueChar("abc")
	if !r {
		t.Error("want true")
	}
}

func TestUniqueChar2(t *testing.T) {
	r := UniqueChar("abcdea")
	if r {
		t.Error("expect false")
	}
}
