package make_two_anagrams

// Make1
// How many steps to make s2 to be an anagram with s1
// https://leetcode.com/problems/minimum-number-of-steps-to-make-two-strings-anagram
func Make1(s1, s2 string) int {
	var (
		m = make(map[rune]int)
	)

	for _, r := range s1 {
		m[r] += 1
	}

	for _, r := range s2 {
		if _, ok := m[r]; ok {
			m[r] -= 1
			if m[r] == 0 {
				delete(m, r)
			}
		}
	}

	var (
		ans int
	)
	for _, v := range m {
		ans += v
	}
	return ans
}
