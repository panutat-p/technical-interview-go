package sqrt

func Get1(num int) int {
	var (
		left  int
		right = num
	)

	for left <= right {
		mid := (left + right) / 2
		if mid*mid > num {
			right = mid - 1
		} else if mid*mid < num {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
