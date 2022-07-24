package array

func LargestSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var best int
	var curr int

	for _, v := range arr {
		curr = Max(v, curr+v)  // drop minus portion from front
		best = Max(best, curr) // drop minus portion from back
	}

	return best
}

func Max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
