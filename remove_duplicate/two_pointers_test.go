package remove_duplicate

import (
	"reflect"
	"testing"
)

func TestRemove1_simple(t *testing.T) {
	want := []int{1, 2}
	got := Remove1([]int{1, 2, 2})
	if !reflect.DeepEqual(want, got) {
		t.Error("want", want, "but got", got)
	}
}

func TestRemove1_swap_with_the_last(t *testing.T) {
	want := []int{2, 11}
	got := Remove1([]int{2, 2, 2, 11})
	if !reflect.DeepEqual(want, got) {
		t.Error("want", want, "but got", got)
	}
}

func TestRemove1_remove_three(t *testing.T) {
	want := []int{2, 3, 6, 9}
	got := Remove1([]int{2, 3, 3, 3, 6, 9, 9})
	if !reflect.DeepEqual(want, got) {
		t.Error("want", want, "but got", got)
	}
}
