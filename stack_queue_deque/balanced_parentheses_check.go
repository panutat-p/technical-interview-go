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

func (q *Stack) Pop() string {
	if len(q.sl) == 1 {
		r := q.sl[0]
		q.sl = []string{}
		return r
	}
	r := q.sl[0]
	q.sl = q.sl[1:len(q.sl)]
	return r
}

func IsBalance(input string) bool {
	pair := map[string]string{"(": ")", "[": "]", "{": "}"}
	keys := maps.Keys(pair)
	stack := Stack{}
	for _, v := range strings.Split(input, "") {
		if slices.Contains(keys, v) {
			stack.Push(pair[v])
		} else if v == stack.Pop() {
			// remove
		} else {
			return false
		}
	}
	return true
}
