package palindrome

import (
	"unicode"
)

// IsPalindrome
// https://leetcode.com/problems/valid-palindrome
func IsPalindrome(input string) bool {
	sl := []rune(input)
	var (
		i = 0
		j = len(sl) - 1
	)
	for i < j {
		if !unicode.IsLetter(sl[i]) {
			i += 1
			continue
		}
		if !unicode.IsLetter(sl[j]) {
			j -= 1
			continue
		}
		if unicode.ToLower(sl[i]) != unicode.ToLower(sl[j]) {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
