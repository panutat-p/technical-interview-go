package main

import (
	"fmt"
)

func main() {
	sl := []int{0, 1, 2, 3}
	fmt.Println("sl: ", sl)

	var (
		sl1 = make([]int, len(sl))
		sl2 = make([]int, len(sl))
	)

	copy(sl1, sl)
	sl1 = sl1[:1]
	fmt.Println("sl1:", sl1)

	copy(sl2, sl)
	sl2 = sl2[2:]
	fmt.Println("sl2:", sl2)
	sl1 = append(sl1, sl2...)
	fmt.Println("sl1:", sl1)
	sl1 = append(sl1, sl2...)
	fmt.Println("sl1:", sl1)
	fmt.Println("sl: ", sl)
}
