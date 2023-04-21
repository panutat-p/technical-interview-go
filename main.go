package main

import "fmt"

func main() {
	sl := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("sl: %v at %p\n", sl, sl)

	sl2 := pop(sl, 1)
	fmt.Printf("sl: %v at %p\n", sl, sl)
	fmt.Printf("sl2: %v at %p\n", sl2, sl2)
}

func insert(nums []int, e int, idx int) []int {
	left := nums[:idx]
	right := append([]int{e}, nums[idx:]...)
	return append(left, right...)
}

func prepend(nums []int, e int) []int {
	return append([]int{e}, nums...)
}

func pop(nums []int, idx int) []int {
	var (
		result = make([]int, 0, len(nums)-1)
	)
	result = append(result, nums[:idx]...)
	return append(result, nums[idx+1:]...)
}
