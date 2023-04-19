package main

import "fmt"

func main() {
	PrintPermutation([]rune{'A', 'B', 'C'}, 0, 3)
}

// PrintPermutation
// https://www.geeksforgeeks.org/write-a-c-program-to-print-all-permutations-of-a-given-string
func PrintPermutation(input []rune, a, b int) {
	if a == b {
		fmt.Println(string(input))
		return
	}

	for i := a; i < b; i += 1 {
		input[i], input[a] = input[a], input[i]
		PrintPermutation(input, a+1, b)
		input[i], input[a] = input[a], input[i]
	}
}
