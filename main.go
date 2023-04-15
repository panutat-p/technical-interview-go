package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Reverse("hello"))
	fmt.Println(Reverse("🐵 eat 🍌"))
}

func Reverse(str string) string {
	var (
		sb strings.Builder
	)

	rsl := []rune(str)
	for i := len(rsl) - 1; i > -1; i -= 1 {
		sb.WriteRune(rsl[i])
	}

	return sb.String()
}
