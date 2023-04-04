package other

import "testing"

func TestFactorial(t *testing.T) {
	got := Factorial(0)
	want := 1
	if got != want {
		t.Error("want:", want, "but got:", got)
	}
}

func TestFactorial2(t *testing.T) {
	got := Factorial(4)
	want := 24
	if got != want {
		t.Error("want:", want, "but got:", got)
	}
}

func TestFactorial3(t *testing.T) {
	got := Factorial(6)
	want := 720
	if got != want {
		t.Error("want:", want, "but got:", got)
	}
}
