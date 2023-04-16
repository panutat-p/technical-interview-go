package shortest_word_distance

func Shortest(words []string, word1, word2 string) int {
	var (
		a   = -1
		b   = -1
		ans = len(words) // the longest distance is len(words) - 1
	)
	for i, e := range words {
		if e == word1 {
			a = i
		} else if e == word2 {
			b = i
		}

		if a != -1 && b != -1 {
			temp := Diff(a, b)
			ans = Min(ans, temp)
		}
	}

	return ans
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
