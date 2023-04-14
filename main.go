package main

import (
	"fmt"
)

func main() {
	fmt.Println("want 12 21 got", permutation("12"))
	fmt.Println("want 123 132 213 231 312 321 got", permutation("123"))
}

func permutation(word string) []string {
	if len(word) == 0 {
		return []string{}
	}

	if len(word) == 1 {
		return []string{word}
	}

	var (
		ans []string
	)

	for i, i32 := range word {
		c := string(i32)
		rest := pop(word, i)

		for _, e := range permutation(rest) {
			ans = append(ans, c+e)
		}
	}
	return ans
}

func pop(str string, idx int) string {
	return str[:idx] + str[idx+1:]
}
