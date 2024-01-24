package word_pattern

import (
	"strings"
)

// TwoMaps
// 290. Word Pattern
// https://leetcode.com/problems/word-pattern
func TwoMaps(patternString, wordString string) bool {
	var (
		m1      = make(map[rune]string) // character
		m2      = make(map[string]rune) // word
		pattern = []rune(patternString)
		words   = strings.Split(wordString, " ")
	)

	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(pattern); i += 1 {
		c := pattern[i]
		w := words[i]
		if v, ok := m1[c]; ok && w != v {
			return false
		}
		if v, ok := m2[w]; ok && c != v {
			return false
		}
		m1[c] = w
		m2[w] = c
	}

	return true
}
