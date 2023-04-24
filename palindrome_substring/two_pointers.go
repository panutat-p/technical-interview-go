package palindrome_substring

// Longest1
// https://leetcode.com/problems/longest-palindromic-substring
func Longest1(input string) string {
	if len(input) <= 1 {
		return input
	}

	sl := []rune(input)
	var (
		ans []rune
	)

	for i := 0; i < len(sl); i += 1 {
		sub := oddPalindrome(sl, i)
		if len(ans) < len(sub) {
			ans = sub
		}
	}

	return string(ans)
}

func oddPalindrome(sl []rune, idx int) []rune {
	var (
		// start from the middle
		i     = idx
		j     = idx
		left  = idx
		right = idx
	)
	for i > -1 && j < len(sl) { // first char and last char cannot go in the loop
		if sl[i] == sl[j] {
			left = i
			right = j
		} else {
			break
		}
		i -= 1 // left shift
		j += 1 // right shift
	}
	return sl[left : right+1]
}
