package string_compression

import "testing"

func TestCompress1_short(t *testing.T) {
	want := "A4B3"
	got := Compress1("AAAABBB")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestCompress1_long(t *testing.T) {
	want := "A4B4C5D2E4"
	got := Compress2("AAAABBBBCCCCCDDEEEE")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
