package remove_duplicate

// Remove1
// Given an longest_parking_roof of sorted numbers, remove all duplicate number instances from it in-place
// The relative order of the elements should be kept the same
// Should not use any extra space
func Remove1(nums []int) []int {
	var (
		idx = 1 // point to non-duplicated element
	)
	for i := 0; i < len(nums); i += 1 {
		if nums[i] != nums[idx-1] {
			nums[idx] = nums[i] // replace the duplicated element
			idx += 1            // right shift pointer
		}
	}

	return nums[:idx]
}
