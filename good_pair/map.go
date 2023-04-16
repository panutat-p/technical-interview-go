package good_pair

import "fmt"

// Count1
// A pair (i, j) is called good if nums[i] == nums[j] and i < j
func Count1(nums []int) int {
	var (
		m   = make(map[int]int)
		ans int
	)

	for _, e := range nums {
		ans += m[e]
		m[e] += 1
	}

	fmt.Println(m)
	return ans
}
