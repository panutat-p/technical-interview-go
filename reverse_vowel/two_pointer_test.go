package reverse_vowel

import "testing"

func TestGenerate2_two_hello(t *testing.T) {
	want := "holle"
	got := Reverse2("hello")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate2_uppercase(t *testing.T) {
	want := "UOIEA"
	got := Reverse2("AEIOU")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate2_mixed(t *testing.T) {
	want := "DusUgnGires"
	got := Reverse2("DesignGUrus")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate2_space(t *testing.T) {
	want := "e crozy mankAy"
	got := Reverse2("A crazy monkey")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
