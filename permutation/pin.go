package permutation

// Generate1
// sorted order
// https://www.geeksforgeeks.org/generate-all-the-permutation-of-a-list-in-python/
func Generate1(perm []string) [][]string {
	if len(perm) == 0 {
		return [][]string{}
	}

	if len(perm) == 1 {
		return [][]string{perm}
	}

	var (
		ans [][]string
	)

	for i := range perm {
		m := perm[i]

		var (
			tmp []string
		)

		tmp = append(tmp, perm[:i]...)
		tmp = append(tmp, perm[i+1:]...)

		for _, v := range Generate1(tmp) {
			var (
				sub []string
			)
			sub = append(sub, m)
			sub = append(sub, v...)

			ans = append(ans, sub)
		}
	}
	return ans
}
