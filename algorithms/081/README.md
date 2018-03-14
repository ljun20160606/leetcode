# Search in Rotated Sorted Array II

## 描述

```txt

Follow up for &#34;Search in Rotated Sorted Array&#34;:
What if duplicates are allowed?

Would this affect the run-time complexity? How and why?


Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).

Write a function to determine if a given target is in the array.

The array may contain duplicates.
```

## 题解

```go
package algorithms

//Follow up for "Search in Rotated Sorted Array":
//What if duplicates are allowed?
//
//Would this affect the run-time complexity? How and why?
//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
//
//Write a function to determine if a given target is in the array.
//
//The array may contain duplicates.

// 上一题 33题
//     0 1 3 3 3 6 7
//           |
// 左侧旋转  |
//     3 3 6 7 0 1 3
// 左侧有序   |
//     3 < 7 > 3
// end <= start < mid

//     0 1 3 3 3 6 7
//              |   右侧旋转
//     6 7 0 1 3 3 3
//           |      右侧有序
//     6 > 1 < 3
// mid <= end < start

func search(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for {
		if start > end {
			break
		}
		mid := (start + end) / 2
		if nums[mid] == target {
			return true
		}
		if nums[start] < nums[mid] { // 左侧旋转
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < nums[end] { // 右侧旋转
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else if nums[start] == nums[mid] {
			start++
		} else {
			end--
		}
	}
	return false
}

```