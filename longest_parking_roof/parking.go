package longest_parking_roof

/*
https://gist.github.com/Rutvik17/d64a85b77a2532dcfffc42657561790a
https://leetcode.com/discuss/interview-question/1317796/amazon-oa-2021-hackerrank
https://www.geeksforgeeks.org/window-sliding-technique
*/

import (
	"fmt"
	"sort"

	"github.com/panutat-p/technical-interview-go/pkg"
)

// ParkingRoof
// alternative, declare best := math.MaxInt, start loop with i:=0
func ParkingRoof(cars []int, k int) int {
	fmt.Println("k:", k)
	sort.Ints(cars)
	fmt.Println("cars:", cars)
	// 0 1 2 3 4
	// o o o _ _
	// _ o o o _
	// _ _ o o o
	best := cars[k-1] - cars[0] + 1
	for i := 1; i < len(cars)-k+1; i += 1 {
		tmp := cars[i+k-1] - cars[i] + 1
		fmt.Println("tmp:", tmp)
		best = pkg.Min(tmp, best)
	}
	return best
}
