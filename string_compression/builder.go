package string_compression

import (
	"fmt"
	"strings"
)

func Compress2(s string) string {
	sb := strings.Builder{}

	var (
		count = 1
	)
	for i := 1; i <= len(s); i++ {
		if i == len(s) {
			sb.WriteString(fmt.Sprintf("%v%d", string(s[i-1]), count))
		} else if s[i] == s[i-1] {
			count += 1
		} else {
			sb.WriteString(fmt.Sprintf("%v%d", string(s[i-1]), count))
			count = 1
		}
	}
	return sb.String()
}
