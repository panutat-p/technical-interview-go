package pair_sum

import "fmt"

// Sum2
// find a pair in the longest_parking_roof whose sum is equal to the given target
// X+Y = Target, search for Y = Target−X, If not found, insert X so that we can search it for the later
// return: longest_parking_roof of index
func Sum2(nums []int, target int) []int {
	var (
		m = make(map[int]int) // k:v = num: index
	)
	for i, e := range nums {
		diff := target - e
		if _, ok := m[diff]; ok {
			fmt.Printf("index (%v, %v) ", m[diff], i)
			fmt.Printf("value (%v, %v)\n", diff, e)
			return []int{m[diff], i}
		} else {
			// diff does not exist
			// store current element and current index
			m[e] = i
		}
	}
	return []int{}
}
