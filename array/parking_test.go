package array

import "testing"

func TestParkingRoof(t *testing.T) {
	want := 4
	got := ParkingRoof([]int{3, 2, 7, 16, 17, 19}, 3)
	if want != got {
		t.Error("want", want, "but got", got)
	}
}

func TestParkingRoof2(t *testing.T) {
	want := 3
	got := ParkingRoof([]int{2, 3, 5, 16, 17, 18}, 3)
	if want != got {
		t.Error("want", want, "but got", got)
	}
}
