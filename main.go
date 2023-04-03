package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("output", CheckPairSum([]int{1, 1, 3, 7, -3, 1}, 4))                                 // 2
	fmt.Println("output", CheckPairSum([]int{3, 2, 3, 4, -3, 1}, 6))                                 // 2
	fmt.Println("output", CheckPairSum([]int{5, 1, 4, -100, 10, -5}, 5))                             // 2
	fmt.Println("output", CheckPairSum([]int{1, 9, 2, 8, 3, 7, 4, 6, 5, 5, 13, 14, 11, 13, -1}, 10)) // 6
}

func CheckPairSum(sl []int, num int) int {
	if len(sl) < 2 {
		return -1
	}

	count := 0
	s := Set{}
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

type Set map[int]struct{}

func (s Set) Add(num int) {
	s[num] = struct{}{}
}

func (s Set) Remove(num int) {
	delete(s, num)
}

func (s Set) IsContain(num int) bool {
	_, ok := s[num]
	return ok
}

func (s Set) Print() {
	if len(s) == 0 {
		fmt.Println("set is empty")
		return
	}

	sb := strings.Builder{}
	for k, _ := range s {
		sb.WriteString(fmt.Sprintf("%v, ", k))
	}
	fmt.Println(sb.String())
}
