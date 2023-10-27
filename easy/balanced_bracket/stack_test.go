package balanced_bracket

import (
	"testing"
)

func TestIsBalance(t *testing.T) {
	yes := IsBalance("[{{{(())}}}]((()))")
	if !yes {
		t.Error("want true")
	}
}

func TestIsBalance2(t *testing.T) {
	no := IsBalance("{[)))")
	if no {
		t.Error("want false")
	}
}

func TestIsBalance3(t *testing.T) {
	no := IsBalance("[[[]])]")
	if no {
		t.Error("want false")
	}
}

func TestIsBalance4(t *testing.T) {
	no := IsBalance("{[]})")
	if no {
		t.Error("want false")
	}
}
