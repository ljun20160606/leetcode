# Search Insert Position

## 描述

```txt
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:
Input: [1,3,5,6], 5
Output: 2



Example 2:
Input: [1,3,5,6], 2
Output: 1



Example 3:
Input: [1,3,5,6], 7
Output: 4



Example 1:
Input: [1,3,5,6], 0
Output: 0


```

## 题解

```go
package algorithms

//Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
//You may assume no duplicates in the array.
//
//Here are few examples.
//[1,3,5,6], 5 → 2
//[1,3,5,6], 2 → 1
//[1,3,5,6], 7 → 4
//[1,3,5,6], 0 → 0

// end = 0
// start = len(nums)-1
//
func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	var mid int
	for start <= end {
		mid = (start + end + 1) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

```