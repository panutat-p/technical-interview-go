package reverse_vowel

func Reverse2(str string) string {
	sl := []rune(str)

	var (
		i = 0
		j = len(sl) - 1
	)
	for i < j {
		// increment until a vowel is found
		for i < j && !isVowel(sl[i]) {
			i += 1
		}

		// decrement until a vowel is found
		for i < j && !isVowel(sl[j]) {
			j -= 1
		}

		// swap the values if both are vowels
		sl[i], sl[j] = sl[j], sl[i]
		i += 1
		j -= 1
	}

	return string(sl)
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
