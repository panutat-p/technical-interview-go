package longest_palindrome

import "unicode"

// Longest2
// https://leetcode.com/problems/longest-palindrome
func Longest2(str string) int {
	if len(str) < 2 {
		return len(str)
	}

	var (
		alphabet [52]int
		ans      int
	)
	for _, r := range str {
		var (
			idx int32
		)
		if unicode.IsUpper(r) {
			idx = r - 'A' + 26
		} else {
			idx = r - 'a'
		}
		alphabet[idx] += 1
		if alphabet[idx]%2 == 0 {
			ans += 2 // even number of characters
		}
	}
	return min(len(str), ans+1) // handle odd number of characters
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
