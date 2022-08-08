package stack_queue_deque

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"strings"
)

type Stack struct {
	sl []string
}

func (q *Stack) Push(s string) {
	q.sl = append([]string{s}, q.sl...)
}

// Pop panic when q.sl is empty
func (q *Stack) Pop() string {
	el := q.sl[0]
	q.sl = q.sl[1:len(q.sl)] // edge case: [100][1:] >> []
	return el
}

func IsBalance(input string) bool {
	pair := map[string]string{"(": ")", "[": "]", "{": "}"}
	keys := maps.Keys(pair)
	stack := Stack{}
	for _, v := range strings.Split(input, "") {
		if slices.Contains(keys, v) {
			stack.Push(pair[v])
		} else if len(stack.sl) != 0 && v == stack.Pop() {
			// remove
		} else {
			return false
		}
	}
	return len(stack.sl) == 0
}
