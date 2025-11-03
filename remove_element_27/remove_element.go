package remove_element_27

func RemoveElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	endIndex := len(nums) - 1
	startIndex := 0

	for {
		if startIndex > endIndex {
			break
		}

		if nums[endIndex] == val {
			endIndex--
			continue
		}

		if nums[startIndex] == val {
			temp := nums[endIndex]
			nums[endIndex] = nums[startIndex]
			nums[startIndex] = temp
			endIndex--
			startIndex++
			continue
		}

		startIndex++
	}

	return startIndex
}
