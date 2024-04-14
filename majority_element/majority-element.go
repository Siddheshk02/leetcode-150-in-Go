package majority_element

func majorityElement(nums []int) int {
	frequencies := make(map[int]int)

	for _, number := range nums {
		frequencies[number]++
	}

	for number, freq := range frequencies {
		if freq > len(nums)/2 {
			return number
		}
	}
	return 0
}
