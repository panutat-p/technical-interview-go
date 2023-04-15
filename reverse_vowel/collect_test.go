package reverse_vowel

import (
	"testing"
)

func TestGenerate1_two_hello(t *testing.T) {
	want := "holle"
	got := Reverse1("hello")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate1_uppercase(t *testing.T) {
	want := "UOIEA"
	got := Reverse1("AEIOU")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate1_mixed(t *testing.T) {
	want := "DusUgnGires"
	got := Reverse1("DesignGUrus")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate1_space(t *testing.T) {
	want := "e crozy mankAy"
	got := Reverse1("A crazy monkey")
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
