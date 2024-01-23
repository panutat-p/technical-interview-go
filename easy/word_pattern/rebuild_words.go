package word_pattern

import (
	"fmt"
	"strings"
)

// RebuildWords
// 290. Word Pattern
// https://leetcode.com/problems/word-pattern
func RebuildWords(pattern, input string) bool {
	letters := []rune(pattern)
	words := strings.Split(input, " ")
	if len(pattern) != len(words) {
		return false
	}

	var (
		m = make(map[rune]string)
	)
	for i := 0; i < len(letters); i += 1 {
		w := words[i]
		c := letters[i]
		m[c] = w
	}
	fmt.Println(m)

	var ans []string
	for i := 0; i < len(letters); i++ {
		c := letters[i]
		w, ok := m[c]
		if !ok {
			return false
		}

		ans = append(ans, w)
	}

	fmt.Println(words, ans)

	return input == strings.Join(ans, " ")
}
