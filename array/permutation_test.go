package array

import (
	"reflect"
	"testing"
)

func TestGeneratePermutation(t *testing.T) {
	want := []string{
		"123",
		"132",
		"213",
		"231",
		"321",
		"312",
	}
	got := GeneratePermutation("123", 0, []string{})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}
