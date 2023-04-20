package string_compression

import (
	"strconv"
	"strings"
)

func Compress1(input string) string {
	var (
		m  = make(map[rune]int)
		sb strings.Builder
	)
	for _, r := range input {
		m[r] += 1
	}

	for k, v := range m {
		sb.WriteRune(k)
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}
