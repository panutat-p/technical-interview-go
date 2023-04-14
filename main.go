package main

import (
	"fmt"
)

func main() {
	fmt.Println("want 5 got", getRoof([]int{12, 6, 5, 2}, 3))
}

// slot with number 1 to n
// k: number of car which cover under the roof
// return minimum roof length
// o o o o o
// 1 2 3 4 5
//
//	[   ]
//
// len = 5, k = 3, want 3 loop
// n = 3 - 1 + 1 = 2
func getRoof(cars []int, k int) int {
	selectionSort(cars)
	fmt.Println(cars)
	ans := cars[k-1] - cars[0] + 1
	fmt.Println("ans", ans)
	for i := 1; i < len(cars)-k+1; i += 1 {
		tmp := cars[k+i-1] - cars[i] + 1
		ans = min(ans, tmp)
		fmt.Println("ans", ans)
	}

	return ans
}

func selectionSort(cars []int) {
	for i := range cars {
		for j := i + 1; j < len(cars); j += 1 {
			if cars[i] > cars[j] {
				cars[i], cars[j] = cars[j], cars[i]
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
