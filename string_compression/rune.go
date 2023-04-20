package string_compression

import "fmt"

func Compress1(input string) string {
	var (
		m   = make(map[rune]int)
		ans string
	)
	for _, r := range input {
		if _, ok := m[r]; ok {
			m[r] += 1
		} else {
			m[r] = 1
		}
	}

	for k, v := range m {
		ans += fmt.Sprintf("%s%d", string(k), v)
		fmt.Println(string(k), v)
	}
	return ans
}