package leetcode

func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	i := 0
	j := 0

	len1 := len(nums1)
	len2 := len(nums2)

	finalNums := make([]int, len1 + len2)

	k := 0

	for i < len1 && j < len2 {
		if nums1[i] <= nums2[j] {
			finalNums[k] = nums1[i]
			i++
		} else {
			finalNums[k] = nums2[j]
			j++
		}
		k++
	}

	for i < len1 {
		finalNums[k] = nums1[i]
		i++
		k++
	}

	for j < len2 {
		finalNums[k] = nums2[j]
		k++
		j++
	}

	medianNumber := k / 2
	if k % 2 == 1 {
		return float64(finalNums[medianNumber])
	} else {
		return float64(finalNums[medianNumber] + finalNums[medianNumber - 1]) / 2
	}
}