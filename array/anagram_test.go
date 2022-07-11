package array

import "testing"

func TestIsAnagram(t *testing.T) {
	if !IsAnagram("dog", "god") {
		t.Error("want true, dog and god")
	}
}

func TestIsAnagram2(t *testing.T) {
	if IsAnagram("fish", "pirate") {
		t.Error("want false, fish and pirate")
	}
}

func TestIsAnagram3(t *testing.T) {
	if IsAnagram("aa", "bb") {
		t.Error("want false, aa and bb")
	}
}

func TestIsAnagram4(t *testing.T) {
	if !IsAnagram("clint eastwood", "old west action") {
		t.Error("want true, clint eastwood and old west action")
	}
}

func TestIsAnagram5(t *testing.T) {
	if !IsAnagram("sweep the floor", "too few helpers") {
		t.Error("want true, sweep the floor and too few helpers")
	}
}
