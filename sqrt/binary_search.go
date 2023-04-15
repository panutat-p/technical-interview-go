package sqrt

// Get1 ^
func Get1(num int) int {
	var (
		left  int
		right = num
	)

	for left <= right {
		mid := (left + right) / 2
		if mid*mid > num {
			// too high
			right = mid - 1
		} else if mid*mid < num {
			// too low
			left = mid + 1
		} else {
			// found
			return mid
		}
	}

	// left > right
	return right
}
