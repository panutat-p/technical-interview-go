package word_pattern

import (
	"strings"
)

func Check1(pattern, input string) bool {
	var (
		set   = make(map[rune]string)
		sl    = []rune(pattern)
		words = strings.Split(input, " ")
	)

	if len(sl) != len(words) {
		return false
	}

	for i := 0; i < len(sl); i += 1 {
		r := sl[i]
		word := words[i]
		v, ok := set[r]
		if ok {
			if v != word {
				return false
			}
		} else {
			set[r] = word
		}
	}

	return true
}
