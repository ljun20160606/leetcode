# Merge Sorted Array

## 描述

```txt
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.


Note:
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2. The number of elements initialized in nums1 and nums2 are m and n respectively.
```

## 题解

```go
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

```