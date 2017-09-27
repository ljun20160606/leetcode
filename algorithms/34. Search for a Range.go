package algorithms

//Given an array of integers sorted in ascending order, find the starting and ending position of a given target value.
//
//Your algorithm's runtime complexity must be in the order of O(log n).
//
//If the target is not found in the array, return [-1, -1].
//
//For example,
//Given [5, 7, 7, 8, 8, 10] and target value 8,
//return [3, 4].

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) <= 0 {
		return result
	}

	start, high := 0, len(nums)-1
	for start < high {
		mid := (start + high) / 2
		if nums[mid] < target {
			start = mid + 1
		} else {
			high = mid
		}
	}

	if nums[start] != target {
		return result
	}

	result[0] = start

	high = len(nums) - 1
	for start < high {
		mid := (start + high + 1) / 2
		if nums[mid] > target {
			high = mid - 1
		} else {
			start = mid
		}
	}
	result[1] = high

	return result
}
