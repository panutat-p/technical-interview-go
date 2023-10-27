package longest_parking_roof

/*
https://gist.github.com/Rutvik17/d64a85b77a2532dcfffc42657561790a
https://www.geeksforgeeks.org/window-sliding-technique
*/

import (
	"fmt"
	"sort"

	"golang.org/x/exp/constraints"
)

// ParkingRoof
// Consider the car parking lot as zero indexed array.
// Given a list of positions occupied by cars.
// Find the min length of roof required to cover k cars.
// https://leetcode.com/discuss/interview-question/1354480/amazon-oa
func ParkingRoof(cars []int, k int) int {
	sort.Ints(cars)
	// 0 1 2 3 4
	// o o o _ _
	// _ o o o _
	// _ _ o o o
	best := cars[k-1] - cars[0] + 1
	for i := 1; i < len(cars)-k+1; i += 1 {
		tmp := cars[i+k-1] - cars[i] + 1
		fmt.Println("tmp:", tmp)
		best = Min(tmp, best)
	}
	return best
}

// alternative, declare best := math.MaxInt, start loop with i:=0

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
