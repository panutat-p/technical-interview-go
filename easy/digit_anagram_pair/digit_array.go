package digit_anagram_pair

func DigitArray(nums []int) int {
	var (
		m = make(map[[10]int]int)
	)
	for _, e := range nums {
		key := count(e)
		m[key] += 1
	}

	var (
		ans int
	)
	for _, v := range m {
		ans += v * (v - 1) / 2
	}
	return ans
}

func count(num int) [10]int {
	var (
		ret [10]int
	)
	for {
		digit := num % 10
		ret[digit] += 1
		num /= 10
		if num == 0 {
			break
		}
	}
	return ret
}
