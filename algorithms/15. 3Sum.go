package algorithms

import (
	"sort"
)

//Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.
//
//Note: The algorithms set must not contain duplicate triplets.
//
//For example, given array S = [-1, 0, 1, 2, -1, -4],
//
//A algorithms set is:
//[
//[-1, 0, 1],
//[-1, -1, 2]
//]

// len nums > 3
// 该思考ksum的问题了
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	length := len(nums) - 2
	end := len(nums) - 1
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, end
		for left < right {
			r := nums[i] + nums[left] + nums[right]
			if r < 0 {
				left++
			} else if r > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right {
					if nums[left] == nums[left+1] {
						left++
					} else if nums[right] == nums[right-1] {
						right--
					} else {
						break
					}
				}
				left++
				right--
			}
		}
	}
	return result
}
