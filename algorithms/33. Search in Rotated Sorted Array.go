package algorithms

//Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
//
//(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
//
//You are given a target value to search. If found in the array return its index, otherwise return -1.
//
//You may assume no duplicate exists in the array.

//     0 1 2 4 5 6 7
//           |
// 左侧旋转  |
//     4 5 6 7 0 1 2
// 左侧有序   |
//     4 < 7 > 2
// end < start < mid

//     0 1 2 4 5 6 7
//              |   右侧旋转
//     6 7 0 1 2 4 5
//           |      右侧有序
//     6 > 1 < 5
// mid < end < start

// iteration
func search(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for {
		if start > end {
			break
		}
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[start] <= nums[mid] { // 左侧旋转
			if target >= nums[start] && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else { // 右侧旋转
			if target > nums[mid] && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
