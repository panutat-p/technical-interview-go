package anagram

import (
	"unicode"
)

func Check1(s1 string, s2 string) bool {
	var (
		m = make(map[rune]int)
	)
	for _, r := range s1 {
		if unicode.IsLetter(r) {
			k := unicode.ToLower(r)
			m[k] += 1
		}
	}

	for _, r := range s2 {
		if unicode.IsLetter(r) {
			k := unicode.ToLower(r)
			m[k] -= 1
			if m[k] == 0 {
				delete(m, k)
			}
		}
	}

	return len(m) == 0
}
