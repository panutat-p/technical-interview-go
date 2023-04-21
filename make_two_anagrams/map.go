package make_two_anagrams

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
