package pair_sum

// Sum2
// find a pair in the array whose sum is equal to the given target
// Let’s say during our iteration we are at number ‘X’, so we need to find ‘Y’ such that “X+Y=target”
// return: indices
func Sum2(nums []int, target int) []int {
	var (
		m = make(map[int]int) // k:v = num: index
	)
	for i, n := range nums {
		if _, ok := m[n]; ok {
			return []int{m[n], i}
		} else {
			// n does not exist
			diff := target - n
			m[diff] = i // store index
		}
	}
	return []int{}
}
