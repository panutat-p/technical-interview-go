package largest_continuous_sum

import (
	"golang.org/x/exp/constraints"
)

// LargestSumKadane
// 53. Maximum Subarray
// https://leetcode.com/problems/maximum-subarray
// Given an integer array nums, find the subarray with the largest sum, and return its sum.
func LargestSumKadane(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		sum int
		ans int
	)
	for _, v := range nums {
		sum = Max(v, sum+v) // drop minus portion from front
		ans = Max(ans, sum) // drop minus portion from back
	}
	return ans
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
