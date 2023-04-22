package group_anagrams

import "fmt"

// Group1
// group the anagrams together
// return: alphabet of group of string in any order
// https://leetcode.com/problems/group-anagrams
func Group1(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}

	if len(words) == 1 {
		return [][]string{words}
	}

	var (
		m   = make(map[[26]int]string)
		ans [][]string
	)
	for _, e := range words {
		alp := alphabet(e)
		if _, ok := m[alp]; ok {
			ans = append(ans, []string{m[alp], e})
			delete(m, alp)
		} else {
			m[alp] = e
		}
	}
	fmt.Println("m", m)
	return ans
}

func alphabet(str string) [26]int {
	var (
		arr [26]int
	)
	for _, r := range str {
		arr[r-'a'] += 1
	}
	return arr
}
