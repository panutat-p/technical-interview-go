package pangram

import (
	"strings"
)

// CheckByMap
// 1832. Check if the Sentence Is Pangram
// https://leetcode.com/problems/check-if-the-sentence-is-pangram
func CheckByMap(input string) bool {
	if len(input) < 26 {
		return false
	}
	var (
		m = make(map[rune]int)
	)
	for _, r := range strings.ToLower(input) {
		if r >= 97 && r <= 122 {
			m[r] += 1
		}
	}
	return len(m) == 26
}
