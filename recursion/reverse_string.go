package recursion

func Reverse(s string) string {
	if len(s) == 0 {
		return ""
	}

	if len(s) == 1 {
		return s
	}

	lastIdx := len(s) - 1
	lastChar := string(s[lastIdx])
	return lastChar + Reverse(s[:lastIdx])
}
