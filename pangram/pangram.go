package pangram

import (
	"strings"
)

func IsPangram(word string) bool {
	var (
		m = make(map[string]int)
	)

	for _, i32 := range strings.ToLower(word) {
		if i32 >= 97 && i32 <= 122 {
			m[string(i32)] += 1
		}
	}

	return len(m) == 26
}
