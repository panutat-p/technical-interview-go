package main

import "fmt"

func main() {
	fmt.Println("want 0 got", Longest(""))
	fmt.Println("want 4 got", Longest("aabb"))
	fmt.Println("want 7 got", Longest("abccccdd"))
}

func Longest(str string) int {
	if len(str) < 2 {
		return len(str)
	}

	var (
		m   = make(map[rune]int)
		ans int
	)

	for _, r := range str {
		m[r] += 1
		if m[r]%2 == 0 {
			ans += 2
		}
	}

	if len(str) == ans {
		return ans
	}

	return ans + 1 // a single character can be placed at the center
}
