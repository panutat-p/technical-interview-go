package sqrt

import "testing"

func TestSqrt_1(t *testing.T) {
	want := 1
	got := Sqrt(1)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestSqrt_2(t *testing.T) {
	want := 2
	got := Sqrt(4)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestSqrt_9(t *testing.T) {
	want := 3
	got := Sqrt(9)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestSqrt_16(t *testing.T) {
	want := 4
	got := Sqrt(16)
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
