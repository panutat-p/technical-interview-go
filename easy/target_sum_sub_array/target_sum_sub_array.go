package target_sum_sub_array

import "fmt"

// SumOfSubArray
// arr = [1,2,3,4,5], target = 9		true
// arr = [1,2,3,4,5], target = 8		false
// arr = [11,9,7,5,3,2], target = 9		true
// arr = [11,9,7,5,3,2], target = 22	false
func SumOfSubArray(arr []int, n int) bool {
	var sum int
	set := make(map[int]struct{})
	for _, v := range arr {
		fmt.Println(set)
		sum += v
		if _, ok := set[sum-n]; ok {
			fmt.Println("🟩 sum", sum, "- target", n, "= excess", sum-n)
			return true
		}

		set[sum] = struct{}{}
	}
	return false
}
