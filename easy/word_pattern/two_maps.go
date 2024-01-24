package word_pattern

import (
	"strings"
)

// TwoMaps
// 290. Word Pattern
// https://leetcode.com/problems/word-pattern
func TwoMaps(pattern, input string) bool {
	var (
		m1    = make(map[rune]string)
		m2    = make(map[string]rune)
		sl    = []rune(pattern)
		words = strings.Split(input, " ")
	)

	if len(sl) != len(words) {
		return false
	}

	for i := 0; i < len(sl); i += 1 {
		r := sl[i]
		word := words[i]
		m1[r] = word
		m2[word] = r
	}

	for i := 0; i < len(sl); i += 1 {
		r := sl[i]
		word := words[i]
		if m1[r] != word || m2[word] != r {
			return false
		}
	}

	return true
}
