package majority_element_169

func MajorityElement(nums []int) int {
	var numDict = make(map[int]int)

	var (
		maxNum      = 0
		maxNumCount = 0
	)

	for index, num := range nums {
		numDict[num]++

		if index == 0 {
			maxNum = num
			maxNumCount = 1
			continue
		}

		if num == maxNum {
			maxNumCount++
		} else if numDict[num] > maxNumCount {
			maxNum = num
		}
	}

	return maxNum
}
