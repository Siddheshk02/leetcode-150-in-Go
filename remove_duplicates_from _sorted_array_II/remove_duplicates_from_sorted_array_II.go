package remove_duplicates_from_sorted_array_II

func removeDuplicates(nums []int) int {
	c := 1
	temp := 0

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			if temp == 0 {
				nums[c] = nums[i]
				c = c + 1
				temp = temp + 1
			} else {
				continue
			}
		} else {
			nums[c] = nums[i]
			c = c + 1
			temp = 0
		}

	}
	return c

}
