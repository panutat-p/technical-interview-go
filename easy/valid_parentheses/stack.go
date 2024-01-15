package valid_parentheses

// IsValid
// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses
func IsValid(input string) bool {
	stack := make(Stack, 0, len(input))
	sl := []rune(input)
	for _, r := range sl {
		switch r {
		case '(':
			stack.Push(')')
		case '[':
			stack.Push(']')
		case '{':
			stack.Push('}')
		default:
			if len(stack) == 0 {
				return false
			}
			if r != stack.Pop() {
				return false
			}
		}
	}
	return len(stack) == 0
}

type Stack []rune

func (s *Stack) Push(r rune) {
	*s = append(*s, r)
}

func (s *Stack) Pop() rune {
	if len(*s) == 0 {
		panic("empty stack")
	}
	r := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return r
}
