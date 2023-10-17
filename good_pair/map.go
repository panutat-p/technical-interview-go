package good_pair

import "fmt"

// 1512. Number of Good Pairs
// https://leetcode.com/problems/number-of-good-pairs
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

func Count2(nums []int) int {
    var (
        ans int
        m = make(map[int]int)
    )
    for _, e := range nums {
        if _, ok := m[e]; ok {
            ans += m[e]
            m[e] += 1
        } else {
            m[e] = 1
        }
    }
    fmt.Println("m:", m)
    return ans
}
