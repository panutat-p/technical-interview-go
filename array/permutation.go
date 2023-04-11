package array

// GeneratePermutation
// https://www.geeksforgeeks.org/find-the-k-th-permutation-sequence-of-first-n-natural-numbers/
func GeneratePermutation(perm string, idx int, result []string) []string {
	if len(perm) == idx {
		result = append(result, perm)
		return result
	}

	for i := idx; i < len(perm); i += 1 {
		perm = SwapChar(perm, i, idx)
		result = GeneratePermutation(perm, idx+1, result)
		perm = SwapChar(perm, i, idx)
	}
	return result
}

func SwapChar(perm string, i, j int) string {
	r := []rune(perm)
	r[i], r[j] = r[j], r[i]
	return string(r)
}
