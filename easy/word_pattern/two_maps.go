package word_pattern

import (
	"strings"
)

// TwoMaps
// 290. Word Pattern
// https://leetcode.com/problems/word-pattern
func TwoMaps(patternString, input string) bool {
	var (
		m1      = make(map[rune]string)
		m2      = make(map[string]rune)
		pattern = []rune(patternString)
		words   = strings.Split(input, " ")
	)

	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(pattern); i += 1 {
		c := pattern[i]
		w := words[i]
		m1[c] = w
		m2[w] = c
	}

	for i := 0; i < len(pattern); i += 1 {
		c := pattern[i]
		w := words[i]
		if m1[c] != w || m2[w] != c {
			return false
		}
	}

	return true
}
