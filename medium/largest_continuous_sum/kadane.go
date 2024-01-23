package largest_continuous_sum

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
		ret = nums[0]
	)
	for _, e := range nums {
		sum = Max(e, sum+e)
		ret = Max(sum, ret)
	}
	return ret
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
