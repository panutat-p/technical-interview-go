package group_anagrams

// Group1
// group the anagrams together
// return: count of group of string in any order
// https://leetcode.com/problems/group-anagrams
func Group1(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}

	if len(words) == 1 {
		return [][]string{words}
	}

	var (
		m = make(map[[26]int][]string)
	)
	for _, e := range words {
		key := count(e)
		m[key] = append(m[key], e)
	}

	var (
		ans [][]string
	)
	for _, k := range m {
		ans = append(ans, k)
	}
	return ans
}

func count(str string) [26]int {
	var (
		ret [26]int
	)
	for _, r := range str {
		ret[r-'a'] += 1
	}
	return ret
}
