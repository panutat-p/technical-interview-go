package permutation

import "sort"

// Generate2
// unsorted order
// https://www.geeksforgeeks.org/find-the-k-th-permutation-sequence-of-first-n-natural-numbers
func Generate2(base string) []string {
	r := accumulate(base, 0, []string{})
	sort.Strings(r)
	return r
}

func accumulate(input string, idx int, result []string) []string {
	if len(input) == idx {
		result = append(result, input)
		return result
	}

	for i := idx; i < len(input); i += 1 {
		input = swap(input, i, idx)
		result = accumulate(input, idx+1, result)
		input = swap(input, i, idx)
	}
	return result
}

// swap
// 'abc' -> 'acb'
func swap(str string, i, j int) string {
	r := []rune(str)
	r[i], r[j] = r[j], r[i]
	return string(r)
}
