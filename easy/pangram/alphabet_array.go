package pangram

// AlphabetArray
// 1832. Check if the Sentence Is Pangram
// https://leetcode.com/problems/check-if-the-sentence-is-pangram
func AlphabetArray(input string) bool {
	if len(input) < 26 {
		return false
	}
	var (
		alp [26]rune
	)
	for _, e := range input {
		alp[e-'a'] = e
	}
	for _, e := range alp {
		if e == 0 {
			return false
		}
	}
	return true
}
