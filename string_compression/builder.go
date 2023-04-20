package string_compression

import (
	"fmt"
	"strings"
)

func Compress2(input string) string {
	sb := strings.Builder{}

	var (
		count = 1
	)
	for i := 1; i <= len(input); i++ {
		if i == len(input) {
			sb.WriteString(fmt.Sprintf("%v%d", string(input[i-1]), count))
		} else if input[i] == input[i-1] {
			count += 1
		} else {
			sb.WriteString(fmt.Sprintf("%v%d", string(input[i-1]), count))
			count = 1
		}
	}
	return sb.String()
}
