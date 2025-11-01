package merge_sorted_array_88

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	var (
		end1   = m - 1
		end2   = n - 1
		input1 = m + n - 1
	)

	for end1 >= 0 && end2 >= 0 {
		if nums1[end1] > nums2[end2] {
			nums1[input1] = nums1[end1]
			end1--
		} else {
			nums1[input1] = nums2[end2]
			end2--
		}

		input1--
	}

	for end2 >= 0 {
		nums1[input1] = nums2[end2]
		end2--
		input1--
	}
}
