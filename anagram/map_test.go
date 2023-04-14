package anagram

import "testing"

func TestCheck1_true(t *testing.T) {
	if !Check1("dog", "god") {
		t.Error("want true, dog and god")
	}
}

func TestCheck1_false(t *testing.T) {
	if Check1("fish", "pirate") {
		t.Error("want false, fish and pirate")
	}
}

func TestCheck1_len(t *testing.T) {
	if !Check1("clint eastwood", "old west action") {
		t.Error("want true, clint eastwood and old west action")
	}
}

func TestCheck1_words(t *testing.T) {
	if !Check1("sweep the floor", "too few helpers") {
		t.Error("want true, sweep the floor and too few helpers")
	}
}
