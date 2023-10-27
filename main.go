package main

import (
	"fmt"
	"github.com/panutat-p/technical-interview-go/pkg"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}
	sl := pkg.Pop(nums, 3)
	fmt.Println(sl)
	sl = pkg.Pop(sl, 3)
	fmt.Println(sl)
	sl = pkg.Pop(sl, 3)
	fmt.Println(sl)
}
