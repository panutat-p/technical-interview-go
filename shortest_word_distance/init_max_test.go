package shortest_word_distance

import "testing"

func TestShortest_simple(t *testing.T) {
	want := 2
	got := Shortest(
		[]string{"a", "b", "c", "d"},
		"a",
		"c",
	)
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestShortest_duplicated(t *testing.T) {
	want := 1
	got := Shortest(
		[]string{"a", "c", "d", "b", "a"},
		"a",
		"b",
	)
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestShortest_dog_fox(t *testing.T) {
	want := 5
	got := Shortest(
		[]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
		"fox",
		"dog",
	)
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
