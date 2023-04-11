package array

import (
	"reflect"
	"testing"
)

func TestGeneratePermutation(t *testing.T) {
	given := "12"
	want := []string{
		"12",
		"21",
	}
	got := GeneratePermutation(given, 0, []string{})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGeneratePermutation2(t *testing.T) {
	given := "123"
	want := []string{
		"123",
		"132",
		"213",
		"231",
		"321",
		"312",
	}
	got := GeneratePermutation(given, 0, []string{})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}
