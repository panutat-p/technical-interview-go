package main

import "fmt"

func main() {
	FindPairSum([]int{6, 4, 2, 1, 3}, 6)
	FindPairSum([]int{9, 11, 2, 5}, 11)
}

func FindPairSum(nums []int, target int) {
	nums = sort(nums)

	var (
		i int
		j = len(nums) - 1
	)
	for i < j {
		if nums[i]+nums[j] < target {
			i += 1
		} else if nums[i]+nums[j] > target {
			j -= 1
		} else {
			fmt.Printf("index (%v, %v) ", i, j)
			fmt.Printf("value (%v, %v)\n", nums[i], nums[j])
			return
		}
	}
}

func sort(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return nums
	}

	pivot := nums[0]
	var (
		left  []int
		right []int
	)
	for _, e := range nums[1:] {
		if pivot > e {
			left = append(left, e)
		} else {
			right = append(right, e)
		}
	}

	nums = append(sort(left), pivot)
	nums = append(nums, sort(right)...)
	return nums
}
