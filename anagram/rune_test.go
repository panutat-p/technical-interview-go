package anagram

import "testing"

func TestCheck2_true(t *testing.T) {
	if !Check2("dog", "god") {
		t.Error("want true, dog and god")
	}
}

func TestCheck2_false(t *testing.T) {
	if Check2("fish", "pirate") {
		t.Error("want false, fish and pirate")
	}
}

func TestCheck2_len(t *testing.T) {
	if !Check2("clint eastwood", "old west action") {
		t.Error("want true, clint eastwood and old west action")
	}
}

func TestCheck2_words(t *testing.T) {
	if !Check2("sweep the floor", "too few helpers") {
		t.Error("want true, sweep the floor and too few helpers")
	}
}
