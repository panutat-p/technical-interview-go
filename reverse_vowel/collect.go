package reverse_vowel

func Reverse1(str string) string {
	var (
		vowels = make([]rune, 0, 100)
		ans    []rune
	)

	for _, r := range str {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowels = append(vowels, r)
		}
	}

	idx := len(vowels) - 1
	for _, r := range str {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			ans = append(ans, vowels[idx])
			idx -= 1
		default:
			ans = append(ans, r)
		}
	}

	return string(ans)
}
