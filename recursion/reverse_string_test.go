package recursion

import "testing"

func TestReverse(t *testing.T) {
	got := Reverse("hello")
	want := "olleh"
	if got != want {
		t.Error("want:", want, "but got:", got)
	}
}

func TestReverse2(t *testing.T) {
	got := Reverse("Go is the best")
	want := "tseb eht si oG"
	if got != want {
		t.Error("want:", want, "but got:", got)
	}
}
