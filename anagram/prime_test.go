package anagram

import "testing"

func TestCheck2_simple(t *testing.T) {
	no := Check2("ac", "bb")
	if no {
		t.Error("want false")
	}
}

func TestCheck2_true(t *testing.T) {
	yes := Check2("dog", "god")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck2_false(t *testing.T) {
	no := Check2("fish", "pirate")
	if no {
		t.Error("want false")
	}
}

func TestCheck2_len(t *testing.T) {
	yes := Check2("clint eastwood", "old west action")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck2_words(t *testing.T) {
	yes := Check2("sweep the floor", "too few helpers")
	if !yes {
		t.Error("want true")
	}
}
