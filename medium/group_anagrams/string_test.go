package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func assert(t *testing.T, got, want [][]string) {
	if len(want) != len(got) {
		t.Error("want", want, "but got", got)
		return
	}

	sort.SliceStable(got, func(i, j int) bool {
		return len(got[i]) < len(got[j])
	})
	for _, e := range got {
		sort.Strings(e)
	}

	if !reflect.DeepEqual(got, want) {
		t.Error("want", want, "but got", got)
	}
}

func TestGroup1_simple(t *testing.T) {
	want := [][]string{
		{"dog", "god"},
	}
	got := Group1([]string{"god", "dog"})
	assert(t, got, want)
}

func TestGroup1_3_chars(t *testing.T) {
	want := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	got := Group1([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	assert(t, got, want)
}

func TestGroup1_one(t *testing.T) {
	want := [][]string{
		{"a"},
	}
	got := Group1([]string{"a"})
	assert(t, got, want)
}

func TestGroup1_empty(t *testing.T) {
	want := [][]string{
		{""},
	}
	got := Group1([]string{""})
	assert(t, got, want)
}
