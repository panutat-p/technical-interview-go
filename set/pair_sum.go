package set

import "fmt"

func IsPairSum(sl []int, num int) int {
	if len(sl) < 2 {
		return 0
	}

	var count int
	set := MySet{}
	for _, v := range sl {
		diff := num - v
		if set.IsInSet(diff) { // found its pair
			count += 1
			fmt.Printf("(%v, %v)\n", diff, v)
		} else {
			set.AddToSet(v)
		}
	}

	// print
	for k := range set {
		fmt.Print(k, " ")
	}

	return count
}

type MySet map[int]struct{}

func (s *MySet) IsInSet(num int) bool {
	_, ok := (*s)[num]
	return ok
}

func (s *MySet) AddToSet(num int) {
	(*s)[num] = struct{}{}
}
