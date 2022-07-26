package other

import (
	"testing"
)

func TestCompress(t *testing.T) {
	r := Compress("AAAABBB")
	if r != "A4B3" {
		t.Error("want", "A4B3", "but got", r)
	}
}

func TestCompress2(t *testing.T) {
	r := Compress("AAAABBBBCCCCCDDEEEE")
	if r != "A4B4C5D2E4" {
		t.Error("want", "A4B4C5D2E4", "but got", r)
	}
}
