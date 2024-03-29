package compress_string

import (
	"strconv"
	"strings"
)

// Compress
// Run-length encoding (RLE) algorithm
// https://www.geeksforgeeks.org/run-length-encoding
func Compress(input string) string {
	sl := []rune(input)
	char := sl[0]
	count := 1
	var sb strings.Builder
	for i := 1; i < len(sl); i++ {
		if sl[i] == sl[i-1] {
			count += 1
		} else {
			sb.WriteRune(char)
			if count > 1 {
				sb.WriteString(strconv.Itoa(count))
			}
			char = sl[i]
			count = 1
		}
	}
	sb.WriteRune(char)
	if count > 1 {
		sb.WriteString(strconv.Itoa(count))
	}
	return sb.String()
}
