package anagram

import (
	"fmt"
	"strings"
)

func Check1(s1 string, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s1 = strings.ToLower(s1)
	s2 = strings.ReplaceAll(s2, " ", "")
	s2 = strings.ToLower(s2)

	fmt.Println(s1)
	fmt.Println(s2)

	if len(s1) != len(s2) {
		return false
	}

	var (
		m = make(map[string]int)
	)
	for _, i32 := range s1 {
		c := string(i32)
		m[c] += 1
	}

	for _, i32 := range s2 {
		c := string(i32)
		m[c] -= 1
		if m[c] == 0 {
			delete(m, c)
		}
	}
	return len(m) == 0
}
