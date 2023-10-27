package permutation

import (
	"reflect"
	"testing"
)

func TestGenerate2_two_char(t *testing.T) {
	given := "12"
	want := []string{
		"12",
		"21",
	}
	got := Generate2(given)
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate2_three_char(t *testing.T) {
	given := "123"
	want := []string{
		"123",
		"132",
		"213",
		"231",
		"312",
		"321",
	}
	got := Generate2(given)
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}
