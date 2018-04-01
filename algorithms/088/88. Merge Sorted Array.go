package algorithms

func merge(nums1 []int, m int, nums2 []int, n int) {
	index := m + n - 1
	i := m - 1
	j := n - 1
	for ; i >= 0 && j >= 0; index-- {
		if nums1[i] >= nums2[j] {
			nums1[index] = nums1[i]
			i--
		} else {
			nums1[index] = nums2[j]
			j--
		}
	}
	for j >= 0 {
		nums1[index] = nums2[j]
		index--
		j--
	}
}
