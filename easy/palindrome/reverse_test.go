package palindrome

import "testing"

func TestCheck1_true(t *testing.T) {
	yes := Check1("Anna")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck1_false(t *testing.T) {
	no := Check1("race a car")
	if no {
		t.Error("want false")
	}
}

func TestCheck1_comma(t *testing.T) {
	yes := Check1("A man, a plan, a canal, Panama!")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck1_special_character(t *testing.T) {
	yes := Check1("Was it a car or a cat I saw?")
	if !yes {
		t.Error("want true")
	}
}

func TestCheck1_number(t *testing.T) {
	no := Check1("0P")
	if no {
		t.Error("want false")
	}
}

func TestCheck1_empty(t *testing.T) {
	yes := Check1(" ")
	if !yes {
		t.Error("want true")
	}
}
