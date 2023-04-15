package permutation

// Generate1
// sorted order
// https://www.geeksforgeeks.org/generate-all-the-permutation-of-a-list-in-python
// example '123'
// 1️⃣ recursion 👉 '1','23' 👉 '1' + '2' + '3'
// 2️⃣ recursion 👉 '2','13' 👉 '2' + '1' + '3'
func Generate1(word string) []string {
	if len(word) == 0 {
		return []string{}
	}

	if len(word) == 1 {
		return []string{word}
	}

	var (
		ans []string
	)

	for i, i32 := range word {
		c := string(i32)
		rest := pop(word, i)

		for _, e := range Generate1(rest) {
			ans = append(ans, c+e)
		}
	}
	return ans
}

func pop(str string, idx int) string {
	return str[:idx] + str[idx+1:]
}
