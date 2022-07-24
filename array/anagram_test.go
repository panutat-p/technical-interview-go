package array

import "testing"

func TestIsAnagramRune(t *testing.T) {
	if !IsAnagramRune("dog", "god") {
		t.Error("want true, dog and god")
	}
}

func TestIsAnagramRune2(t *testing.T) {
	if IsAnagramRune("fish", "pirate") {
		t.Error("want false, fish and pirate")
	}
}

func TestIsAnagramRune3(t *testing.T) {
	if IsAnagramRune("aa", "bb") {
		t.Error("want false, aa and bb")
	}
}

func TestIsAnagramRune4(t *testing.T) {
	if !IsAnagramRune("clint eastwood", "old west action") {
		t.Error("want true, clint eastwood and old west action")
	}
}

func TestIsAnagramRune5(t *testing.T) {
	if !IsAnagramRune("sweep the floor", "too few helpers") {
		t.Error("want true, sweep the floor and too few helpers")
	}
}

func TestIsAnagramMap(t *testing.T) {
	if !IsAnagramMap("sweep the floor", "too few helpers") {
		t.Error("want true")
	}
}
