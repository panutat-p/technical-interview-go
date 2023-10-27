package string_decompression

import (
	"strconv"
	"strings"
)

func Decode1(s string) string {
	var result strings.Builder
	i := 0

	for i < len(s) {
		if s[i] == '[' {
			count, _ := strconv.Atoi(string(s[i-1]))
			substring := Decode1(s[i+1:])

			result.WriteString(strings.Repeat(substring, count))
			i += len(substring) + 2 // Skip the processed substring and closing bracket
		} else if s[i] == ']' {
			return result.String()
		} else {
			result.WriteByte(s[i])
			i++
		}
	}

	return result.String()
}

func INT(r rune) int {
	return int(r - '0')
}
