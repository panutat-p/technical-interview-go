package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func TestGroup1_simple(t *testing.T) {
	want := [][]string{
		{"dog", "god"},
	}
	got := Group1([]string{"god", "dog"})

	for _, e1 := range got {
		sort.Strings(e1)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGroup1_3_chars(t *testing.T) {
	want := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	got := Group1([]string{"eat", "tea", "tan", "ate", "nat", "bat"})

	for _, e1 := range got {
		sort.Strings(e1)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGroup1_one(t *testing.T) {
	want := [][]string{
		{"a"},
	}
	got := Group1([]string{"a"})

	for _, e1 := range got {
		sort.Strings(e1)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGroup1_empty(t *testing.T) {
	want := [][]string{
		{""},
	}
	got := Group1([]string{""})

	for _, e1 := range got {
		sort.Strings(e1)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}
