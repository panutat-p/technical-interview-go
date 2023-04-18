package main

import "fmt"

func main() {
	//fmt.Println("want 12, 21 got", permutation("12"))
	fmt.Println("want 123, 132, 213, 231, 312, 321 got", permutation("1234"))
}

func permutation(input string) []string {
	if len(input) == 0 {
		return []string{}
	}

	if len(input) == 1 {
		return []string{input}
	}

	var (
		ans []string
	)

	sl := []rune(input)
	for i, r := range input {
		rest := pop(sl, i)
		// tmp := append([]rune{r}, rest...)
		for _, s := range permutation(string(rest)) {
			tmp := string(r) + s
			ans = append(ans, tmp)
		}
	}

	return ans
}

func pop(sl []rune, idx int) []rune {
	var (
		rest = make([]rune, 0, len(sl))
	)
	rest = append(rest, sl[:idx]...)
	rest = append(rest, sl[idx+1:]...)
	return rest
}
