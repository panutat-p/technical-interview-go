package find_missing_values

import "testing"

func TestFind2_missing_five(t *testing.T) {
	want := 5
	got := Find2([]int{5, 5, 5, 5, 7, 7}, []int{5, 5, 5, 7, 7})
	if got != want {
		t.Error("want", want, "but got", got)
	}
}
