package array

import (
	"fmt"
	"strings"
)

// IsAnagramRune
// sum of rune value
func IsAnagramRune(s1 string, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")

	if len(s1) != len(s2) {
		return false
	}

	var (
		total1 int32
		total2 int32
	)

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

func IsAnagramMap(s1 string, s2 string) bool {
	s1 = strings.ReplaceAll(s1, " ", "")
	s2 = strings.ReplaceAll(s2, " ", "")

	if len(s1) != len(s2) {
		return false
	}

	m := make(map[string]int)
	for _, r := range s1 {
		s := string(r)
		if _, ok := m[s]; ok {
			m[s] += 1
		} else {
			m[s] = 1
		}
	}

	for _, r := range s2 {
		s := string(r)
		if _, ok := m[s]; ok {
			m[s] -= 1
			if m[s] == 0 {
				delete(m, s)
			}
		} else {
			return false
		}
	}
	return true
}

type MyMap map[string]int

func (m MyMap) Key(s string) bool {
	_, ok := m[s]
	return ok
}

func (m MyMap) Add(s string) {
	if _, ok := m[s]; ok {
		m[s] += 1
	} else {
		m[s] = 1
	}
}

func (m MyMap) Print() {
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%v: %v", k, v))
	}
	fmt.Println(sb.String())
}
