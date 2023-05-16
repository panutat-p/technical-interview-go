package main

import "fmt"

func main() {
	var (
		sl Slice
	)

	sl.Append(1)
	sl.Append(2)
	sl.Append(3)
	sl.Append(4)
	fmt.Println(len(sl), cap(sl), sl)
	sl.Pop(len(sl) - 1)
	fmt.Println(len(sl), cap(sl), sl)
}

type Slice []int

func (s *Slice) Append(e int) {
	*s = append(*s, e)
}

func (s *Slice) Prepend(e int) {
	*s = append([]int{e}, *s...)
}

func (s *Slice) Insert(e, idx int) {
	left := (*s)[:idx]
	right := (*s)[idx:]
	*s = append(left, e)
	*s = append(*s, right...)
}

func (s *Slice) Pop(idx int) int {
	ret := (*s)[idx]
	*s = append((*s)[:idx], (*s)[idx+1:]...)
	return ret // same backing array
}
