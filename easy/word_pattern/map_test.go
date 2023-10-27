package word_pattern

import "testing"

func TestCheck1_true(t *testing.T) {
	yes := Check1("abba", "dog cat cat dog")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck1_false(t *testing.T) {
	no := Check1("abba", "dog cat cat fish")
	if no {
		t.Error("want false")
	}
}

func TestCheck1_same_char(t *testing.T) {
	no := Check1("aaaa", "dog cat cat dog")
	if no {
		t.Error("want false")
	}
}

func TestCheck1_same_word(t *testing.T) {
	no := Check1("abba", "dog dog dog dog")
	if no {
		t.Error("want false")
	}
}
