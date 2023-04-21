package anagram

import "testing"

func TestCheck1_simple(t *testing.T) {
	no := Check1("ac", "bb")
	if no {
		t.Error("want false")
	}
}

func TestCheck1_true(t *testing.T) {
	yes := Check1("dog", "god")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck1_false(t *testing.T) {
	no := Check1("fish", "pirate")
	if no {
		t.Error("want false")
	}
}

func TestCheck1_len(t *testing.T) {
	yes := Check1("clint eastwood", "old west action")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck1_words(t *testing.T) {
	yes := Check1("sweep the floor", "too few helpers")
	if !yes {
		t.Error("want true")
	}
}
