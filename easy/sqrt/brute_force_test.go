package sqrt

import "testing"

func TestGet2_1(t *testing.T) {
	want := 1
	got := Get2(1)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGet2_2(t *testing.T) {
	want := 2
	got := Get2(4)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGet2_9(t *testing.T) {
	want := 3
	got := Get2(9)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGet2_16(t *testing.T) {
	want := 4
	got := Get2(16)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGet2_rounded_down_to_the_nearest_integer(t *testing.T) {
	want := 1
	got := Get2(2)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
