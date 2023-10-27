package permutation

import (
	"reflect"
	"testing"
)

func TestGenerate1_two_char(t *testing.T) {
	want := []string{
		"12",
		"21",
	}
	got := Generate1("12")
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate1_three_char(t *testing.T) {
	want := []string{
		"123",
		"132",
		"213",
		"231",
		"312",
		"321",
	}
	got := Generate1("123")
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}
