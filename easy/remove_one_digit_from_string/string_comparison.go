package remove_one_digit_from_string

// StringComparison
// Input: number = 123, digit = 3    ->  Output: 12
// Input: number = 1231, digit = 1   ->  Output: 231
// Input: number = 551, digit = 5    ->  Output: 51
// Input: number = -4674, digit = 4  ->  Output: -467
// Input: number = 5552, digit = 5   ->  Output: 552
func StringComparison(num string, inputDigit byte) string {
	var (
		rsl   = []rune(num)
		sl    []string
		digit = rune(inputDigit)
	)
	for i, e := range rsl {
		if e == digit {
			sl = append(sl, string(Pop(rsl, i)))
		}
	}
	var (
		ans = sl[0]
	)
	for _, e := range sl[1:] {
		ans = Max(ans, e)
	}
	return ans
}

func Pop(sl []rune, idx int) []rune {
	var (
		ret = make([]rune, len(sl)-1)
	)
	copy(ret, sl[:idx])
	copy(ret[idx:], sl[idx+1:])
	return ret
}

func Max(a, b string) string {
	if len(a) != len(b) {
		panic("unequal length")
	}
	if a > b {
		return a
	}
	return b
}
