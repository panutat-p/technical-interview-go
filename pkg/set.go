package pkg

import (
	"fmt"
	"strings"
)

type Set map[int]struct{}

func NewSet() Set {
	return Set{}
}

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
