package palindrome_substring

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
