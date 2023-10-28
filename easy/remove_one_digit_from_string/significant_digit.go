package remove_one_digit_from_string

import (
	"slices"
)

func SignificantDigit(inputNumber string, inputDigit byte) string {
	digit := rune(inputDigit)
	sl := []rune(inputNumber)
	var (
		idx int
	)
	for i := 0; i < len(inputNumber); i += 1 {
		if sl[i] == digit {
			idx = i
			if i < len(inputNumber)-1 && sl[i] < sl[i+1] {
				break
			}
		}
	}
	return string(slices.Delete(sl, idx, idx+1))
}
