package _map

import (
	"fmt"
	"math"
)

func FindMissing(arr1 []int, arr2 []int) int {
	var m = make(map[int]int)
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	for _, v := range arr1 {
		if m[v] == 0 {
			return v
		} else {
			m[v] -= 1
		}
	}
	return math.MinInt
}

func FindMissing2(arr1 []int, arr2 []int) int {
	var m = make(map[int]int)
	for _, v := range arr1 {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	fmt.Println(m)
	for _, v := range arr2 {
		if _, ok := m[v]; ok {
			m[v] -= 1
		}
	}
	fmt.Println(m)

	var result []int
	for k, v := range m {
		if v != 0 {
			result = append(result, k)
		}
	}
	fmt.Println(result)
	return result[0]
}

func FindMissingBruteForce(arr1 []int, arr2 []int) int {
	for _, v := range arr1 {
		var count int
		for _, v2 := range arr2 {
			if v == v2 {
				count += 1
			}
		}

		if !isInArray(arr2, v) {
			return v
		}
	}
	return -1
}

func isInArray(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}
