package hashmap

import "math"

func Find2(arr1 []int, arr2 []int) int {
	for _, v := range arr1 {
		idx := index(arr2, v)
		if idx == -1 {
			// not found in arr2
			return v
		} else {
			// found in arr2
			arr2 = pop(arr2, idx)
		}
	}
	return math.MinInt
}

func index(arr []int, num int) int {
	for i, e := range arr {
		if e == num {
			return i
		}
	}
	return -1
}

func pop(nums []int, idx int) []int {
	return append(nums[:idx], nums[idx+1:]...)
}
