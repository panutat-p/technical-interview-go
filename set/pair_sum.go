package set

import (
	"fmt"
	"strings"
)

func IsPairSum(sl []int, num int) int {
	if len(sl) < 2 {
		return 0
	}

	var count int
	set := MySet{}
	for _, v := range sl {
		diff := num - v
		if set.Contains(diff) {
			count += 1
			set.Remove(diff)
			fmt.Printf("(%v, %v)\n", diff, v)
		} else {
			set.Add(v)
		}
	}

	set.Print()

	return count
}

type MySet map[int]struct{}

func (s *MySet) Contains(num int) bool {
	_, ok := (*s)[num]
	return ok
}

func (s *MySet) Add(num int) {
	(*s)[num] = struct{}{}
}

func (s *MySet) Remove(num int) {
	delete(*s, num)
}

func (s *MySet) Print() {
	if len(*s) == 0 {
		fmt.Println("🟧 set is empty")
		return
	}
	sb := strings.Builder{}
	for k := range *s {
		sb.WriteString(fmt.Sprintf("%v, ", k))
	}
	fmt.Println("🟧 set:", sb.String())
}
