package pangram

import (
	"fmt"
	"strings"
)

func IsPangram(word string) bool {
	var (
		m = make(map[rune]int)
	)

	for _, r := range strings.ToLower(word) {
		if r >= 97 && r <= 122 {
			m[r] += 1
		}
	}
	fmt.Println(m)
	return len(m) == 26
}
