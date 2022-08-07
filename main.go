package main

import (
	"fmt"
	"golang.org/x/exp/slices"
)

func main() {
	// Index returns the index of the first occurrence of v in s, or -1 if not present.
	r := slices.Index([]int{2, 3, 1, 5, 3, 5, 3}, 5)
	fmt.Println(r) // 3
}
