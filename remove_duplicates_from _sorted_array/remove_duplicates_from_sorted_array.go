package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	c := 0
	nums[c] = nums[0]
	c = c + 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		} else {
			nums[c] = nums[i]
			c = c + 1
		}
	}

	return c
}
