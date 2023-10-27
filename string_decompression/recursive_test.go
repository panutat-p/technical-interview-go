package string_decompression

import "testing"

func TestDecode1_simple(t *testing.T) {
	want := "xxx"
	got := Decode1("3[x]")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestDecode1_word(t *testing.T) {
	got := "hellohellohello"
	want := Decode1("3[hello]")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
