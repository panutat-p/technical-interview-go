package palindrome

import (
	"unicode"
)

func Check1(str string) bool {
	var (
		original []rune
	)
	// filter only alphabet in lowercase
	for _, r := range str {
		if unicode.IsLetter(r) {
			original = append(original, unicode.ToLower(r))
		}
	}

	var (
		reversed []rune
	)
	// reverse the string
	for i := len(original) - 1; i > -1; i -= 1 {
		reversed = append(reversed, original[i])
	}

	return string(original) == string(reversed)
}
