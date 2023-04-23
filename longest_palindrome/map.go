package longest_palindrome

// Longest1
// Letters are case sensitive, for example, "Aa" is not considered a palindrome
// https://leetcode.com/problems/longest-palindrome
func Longest1(str string) int {
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
