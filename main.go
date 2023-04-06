package main

import (
	"fmt"
)

func main() {
	fmt.Println("🟩 want 5, got", LargestSum([]int{-4, -5, 5}))
	fmt.Println("🟩 want 25, got", LargestSum([]int{1, 2, -1, 3, 10, 10, -10 - 1}))
}

func LargestSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		sum  int
		best int
	)
	for _, v := range nums {
		sum = Max(sum+v, v)
		best = Max(sum, best)
	}

	return best
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
