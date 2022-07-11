package array

import (
	"strings"
)

// IsAnagramRune
// sum of rune value
func IsAnagramRune(s1 string, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")

	var total1 int32
	var total2 int32

	r1 := []rune(s1)
	for _, r := range r1 {
		total1 += r
	}

	r2 := []rune(s2)
	for _, r := range r2 {
		total2 += r
	}
	return total1 == total2
}
