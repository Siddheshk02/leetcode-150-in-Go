package remove_element

func removeElement(nums []int, val int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[c] = nums[i]
		c = c + 1
	}

	return c
}
