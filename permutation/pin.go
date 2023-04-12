package permutation

// Generate1
// sorted order
// https://www.geeksforgeeks.org/generate-all-the-permutation-of-a-list-in-python
// '123' 👉 '1' + '23' 👉 '1' + '2' + '3'
//
//	              👉 '1' + '3' + '2'
//	👉 '2' + '13' 👉 '2' + '1' + '3'
//	              👉 '2' + '3' + '1'
func Generate1(base string) []string {
	if len(base) == 0 {
		return []string{}
	}

	if len(base) == 1 {
		return []string{base}
	}

	var (
		ans []string
	)

	for i := range base {
		pin := base[i : i+1] // pin char at index 0, 1, 2, ...

		tmp := base[:i] + base[i+1:]

		for _, v := range Generate1(tmp) {
			perm := pin + v
			ans = append(ans, perm)
		}
	}
	return ans
}
