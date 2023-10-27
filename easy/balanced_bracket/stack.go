package balanced_bracket

func IsBalance(input string) bool {
	var (
		stack Stack
	)
	for _, r := range input {
		switch r {
		case '(':
			stack.Push(')')
		case '{':
			stack.Push('}')
		case '[':
			stack.Push(']')
		default:
			if len(stack) == 0 {
				return false
			} else {
				if r != stack.Pop() {
					return false
				}
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
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}
