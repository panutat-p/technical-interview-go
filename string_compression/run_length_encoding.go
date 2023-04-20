package string_compression

import (
	"strconv"
	"strings"
)

// Compress2
// https://www.geeksforgeeks.org/run-length-encoding
func Compress2(input string) string {
	sl := []rune(input)

	var (
		sb    strings.Builder
		char  = sl[0]
		count = 1
	)
	for i := 1; i < len(sl); i++ {
		if sl[i] == sl[i-1] {
			count += 1
		} else {
			sb.WriteRune(char)
			sb.WriteString(strconv.Itoa(count))
			char = sl[i]
			count = 1
		}
	}
	sb.WriteRune(char)
	sb.WriteString(strconv.Itoa(count))
	return sb.String()
}
