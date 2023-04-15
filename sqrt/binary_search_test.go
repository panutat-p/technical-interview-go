package sqrt

import "testing"

func TestGet1_1(t *testing.T) {
	want := 1
	got := Get1(1)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGet1_2(t *testing.T) {
	want := 2
	got := Get1(4)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGet1_9(t *testing.T) {
	want := 3
	got := Get1(9)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGet1_16(t *testing.T) {
	want := 4
	got := Get1(16)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
