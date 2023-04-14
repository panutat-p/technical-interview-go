package anagram

import "strings"

func Check2(s1, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s1 = strings.ToLower(s1)
	s2 = strings.ReplaceAll(s2, " ", "")
	s2 = strings.ToLower(s2)

	if len(s1) != len(s2) {
		return false
	}

	var (
		r1 int32
		r2 int32
	)

	for _, v := range s1 {
		r1 += v
	}

	for _, v := range s2 {
		r2 += v
	}

	return r1 == r2
}
