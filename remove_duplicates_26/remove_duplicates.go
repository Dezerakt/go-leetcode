package remove_duplicates_26

func RemoveDuplicates(nums []int) int {
	// correct, but takes a lot of cpu mc
	var (
		lastValue       = 0
		lastSwitchedIdx = 0
	)

	for index, num := range nums {
		if num <= lastValue && index != 0 {

			for i := index; i < len(nums); i++ {
				if nums[i] > lastValue {
					lastSwitchedIdx++
					nums[index] = nums[i]
					lastValue = nums[i]
					break
				}
			}
		} else {
			lastSwitchedIdx++
			lastValue = num
		}
	}

	//////////////////////////////////////////////////

	return lastSwitchedIdx
}

// 0, 0, 1, 1, 1, 2, 2, 3, 3, 4
// 0, 1, 2, 3, 4, 0, 0, 0, 0, 0
