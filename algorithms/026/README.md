# Remove Duplicates from Sorted Array

## 描述

```txt

Given a sorted array, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.


Example:
Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
It doesn&#39;t matter what you leave beyond the new length.


```

## 题解

```go
package algorithms

//Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.
//
//Do not allocate extra space for another array, you must do this in place with constant memory.
//
//For example,
//Given input array nums = [1,1,2],
//
//Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.

func removeDuplicates(nums []int) int {
	length, newLength := len(nums), 1
	for i := 1; i < length; i++ {
		if nums[i-1] != nums[i] {
			nums[newLength] = nums[i]
			newLength++
		}
	}
	return newLength
}

```