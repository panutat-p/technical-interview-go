package hashmap

import (
	"math"
)

func Find1(arr1 []int, arr2 []int) int {
	var m = make(map[int]int)
	for _, e := range arr2 {
		m[e] += 1
	}

	for _, e := range arr1 {
		if m[e] == 0 {
			return e
		} else {
			m[e] -= 1
		}
	}
	return math.MinInt
}
