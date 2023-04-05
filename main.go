package main

import (
	"fmt"

	"github.com/panutat-p/technical-interview-go/pkg"
)

func main() {
	fmt.Println("🟩 want 2, got", CheckPairSum([]int{1, 1, 3, 7, -3, 1}, 4))
	fmt.Println("🟩 want 2, got", CheckPairSum([]int{3, 2, 3, 4, -3, 1}, 6))
	fmt.Println("🟩 want 2, got", CheckPairSum([]int{5, 1, 4, -100, 10, -5}, 5))
	fmt.Println("🟩 want 6, got", CheckPairSum([]int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5, 13, 14, 11, 13, -1}, 10))
}

func CheckPairSum(sl []int, num int) int {
	if len(sl) < 2 {
		return -1
	}

	count := 0
	s := pkg.NewSet()
	for _, v := range sl {
		target := num - v
		if s.IsContain(v) {
			count += 1
		} else {
			s.Add(target)
		}
	}
	return count
}
