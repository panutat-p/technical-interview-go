package hashmap

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

	for k, v := range m {
		if v != 0 {
			return k
		}
	}
	return math.MinInt
}

func FindMissingBruteForce(arr1 []int, arr2 []int) int {
	for _, v := range arr1 {
		idx := find(arr2, v)
		if idx == -1 {
			return v
		} else {
			arr2 = append(arr2[0:idx], arr2[idx+1:]...)
			fmt.Println("arr2: remove", arr2[idx], "at", idx)
		}
	}
	return -1
}

func find(arr []int, num int) int {
	for i, v := range arr {
		if v == num {
			return i
		}
	}
	return -1
}
