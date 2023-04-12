package array

import "github.com/panutat-p/technical-interview-go/pkg"

func LargestSum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	var (
		best int
		curr int
	)

	for _, v := range arr {
		curr = pkg.Max(v, curr+v)  // drop minus portion from front
		best = pkg.Max(best, curr) // drop minus portion from back
	}

	return best
}
