package shortest_word_distance

func Shortest(words []string, word1, word2 string) int {
	var (
		a   = -1
		b   = len(words)
		ans = len(words) // the possible longest distance is len(words) - 1
	)
	for i, e := range words {
		switch e {
		case word1:
			a = i
			ans = min(ans, distance(a, b))
		case word2:
			b = i
			ans = min(ans, distance(a, b))
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func distance(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
