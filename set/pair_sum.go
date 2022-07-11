package set

import "fmt"

func IsPairSum(sl []int, num int) int {
	if len(sl) < 2 {
		return 0
	}

	var count int
	set := make(map[int]bool) // set store diff value
	for _, v := range sl {
		diff := num - v
		if IsInSet(set, diff) { // found its pair
			count += 1
		} else {
			AddToSet(set, v)
		}
	}

	// print
	for k := range set {
		fmt.Print(k, " ")
	}

	return count
}

func IsInSet(set map[int]bool, num int) bool {
	_, ok := set[num]
	return ok
}

func AddToSet(set map[int]bool, num int) {
	set[num] = true
}
