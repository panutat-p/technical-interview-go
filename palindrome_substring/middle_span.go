package palindrome_substring

// Longest1
// https://leetcode.com/problems/longest-palindromic-substring
func Longest1(input string) string {
	if len(input) <= 1 {
		return input
	}

	sl := []rune(input)
	var (
		start int
		end   int
		left  int
		right int
	)

	for i := 0; i < len(sl); i += 1 {
		// result is odd length
		left = i
		right = i

		for left > -1 && right < len(sl) && sl[left] == sl[right] {
			if end-start < right-left {
				start = left
				end = right
			}
			left -= 1
			right += 1
		}

		// result is even length
		left = i
		right = i + 1

		for left > -1 && right < len(sl) && sl[left] == sl[right] {
			if end-start < right-left {
				start = left
				end = right
			}
			left -= 1
			right += 1
		}

	}

	return string(sl[start : end+1])
}
