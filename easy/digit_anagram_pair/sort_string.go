package digit_anagram_pair

import (
	"strconv"
)

func Count1(nums []int) int {
	var (
		m = make(map[string]int)
	)
	for _, e := range nums {
		m[chars(e)] += 1
	}

	var (
		ans int
	)
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	return ans
}

func chars(num int) string {
	s := strconv.Itoa(num)
	var (
		sl []rune
	)
	for _, r := range s {
		sl = append(sl, r)
	}
	sort(sl)
	return string(sl)
}

func sort(rsl []rune) {
	for i := range rsl {
		for j := i + 1; j < len(rsl); j += 1 {
			if rsl[i] > rsl[j] {
				rsl[i], rsl[j] = rsl[j], rsl[i]
			}
		}
	}
}
