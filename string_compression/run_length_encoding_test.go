package string_compression

import (
	"testing"
)

func TestCompress2_short(t *testing.T) {
	r := Compress2("AAAABBB")
	if r != "A4B3" {
		t.Error("want", "A4B3", "but got", r)
	}
}

func TestCompress2_long(t *testing.T) {
	r := Compress2("AAAABBBBCCCCCDDEEEE")
	if r != "A4B4C5D2E4" {
		t.Error("want", "A4B4C5D2E4", "but got", r)
	}
}
