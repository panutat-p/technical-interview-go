package pair_sum

// Sum1
// find a pair in the array whose sum is equal to the given target
// return: indices
func Sum1(nums []int, target int) []int {
	sort(nums)

	var (
		i = 0
		j = len(nums) - 1
	)
	for i < j {
		sum := nums[i] + nums[j]
		if sum > target {
			j -= 1
		} else if sum < target {
			i += 1
		} else {
			return []int{i, j}
		}
	}
	return []int{}
}

func sort(nums []int) {
	for i := 0; i < len(nums); i += 1 {
		for j := i + 1; j < len(nums); j += 1 {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
