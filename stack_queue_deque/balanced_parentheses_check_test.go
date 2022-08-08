package stack_queue_deque

import (
	"testing"
)

func TestIsBalance(t *testing.T) {
	want := true
	got := IsBalance("[{{{(())}}}]((()))")
	if want != got {
		t.Error("want", want)
	}
}

func TestIsBalance2(t *testing.T) {
	want := false
	got := IsBalance("{[)))")
	if want != got {
		t.Error("want", want)
	}
}

func TestIsBalance3(t *testing.T) {
	want := false
	got := IsBalance("[[[]])]")
	if want != got {
		t.Error("want", want)
	}
}

func TestIsBalance4(t *testing.T) {
	want := false
	got := IsBalance("{[]})")
	if want != got {
		t.Error("want", want)
	}
}
