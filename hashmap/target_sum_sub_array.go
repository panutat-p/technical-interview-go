package hashmap

import "fmt"

// SumOfSubArray
// arr = [1,2,3,4,5], target = 9		true
// arr = [1,2,3,4,5], target = 8		false
// arr = [11,9,7,5,3,2], target = 9		true
// arr = [11,9,7,5,3,2], target = 22	false
func SumOfSubArray(arr []int, n int) bool {
	for _, v := range arr {
		if v == n {
			return true
		}
	}

	var sum int
	m := make(map[int]int)
	for _, v := range arr {
		fmt.Println(m)
		sum += v
		if _, ok := m[sum-n]; ok {
			return true
		}

		if _, ok := m[sum]; ok {
			m[sum] += 1
		} else {
			m[sum] = 1
		}
	}
	return false
}
