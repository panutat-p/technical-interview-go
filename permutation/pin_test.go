package permutation

import (
	"reflect"
	"testing"
)

func TestGenerate_two_char(t *testing.T) {
	want := [][]string{
		{"1", "2"},
		{"2", "1"},
	}
	got := Generate1([]string{"1", "2"})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGenerate_three_char(t *testing.T) {
	want := [][]string{
		{"1", "2", "3"},
		{"1", "3", "2"},
		{"2", "1", "3"},
		{"2", "3", "1"},
		{"3", "1", "2"},
		{"3", "2", "1"},
	}
	got := Generate1([]string{"1", "2", "3"})
	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}
