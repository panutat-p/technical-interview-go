package palindrome

import (
	"unicode"
)

// Check2
// Read only alphanumeric characters
// https://leetcode.com/problems/valid-palindrome
func Check2(input string) bool {
	var (
		sl = make([]rune, 0, len(input))
	)
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sl = append(sl, unicode.ToLower(r))
		}
	}

	if len(sl) <= 1 {
		return true
	}

	var (
		i = 0
		j = len(sl) - 1
	)

	for i < j {
		if sl[i] != sl[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
