package largest_continuous_sum

import "github.com/panutat-p/technical-interview-go/pkg"

// 53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray
// Given an integer array nums, find the subarray with the largest sum, and return its sum.

func LargestSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var (
		best int
		curr int
	)

	for _, v := range arr {
		curr = pkg.Max(v, curr+v)  // drop minus portion from front
		best = pkg.Max(best, curr) // drop minus portion from back
	}

	return best
}

func Sum2(nums []int) int {
    var sum int
    var ans = nums[0]
    for _, e := range nums {
        sum = max(sum+e, e)
        ans = max(ans, sum)
    }
    return ans
}

func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
