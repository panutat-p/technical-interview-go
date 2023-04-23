package palindrome

import "testing"

func TestCheck2_true(t *testing.T) {
	yes := Check2("Anna")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck2_false(t *testing.T) {
	no := Check2("race a car")
	if no {
		t.Error("want false")
	}
}

func TestCheck2_comma(t *testing.T) {
	yes := Check2("A man, a plan, a canal, Panama!")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck2_special_character(t *testing.T) {
	yes := Check2("Was it a car or a cat I saw?")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck2_number(t *testing.T) {
	no := Check2("0P")
	if no {
		t.Error("want false")
	}
}

func TestCheck2_empty(t *testing.T) {
	yes := Check2(" ")
	if !yes {
		t.Error("want true")
	}
}
