package algorithms

import (
	"sort"
)

//Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.
//
//Note: The algorithms set must not contain duplicate quadruplets.
//
//For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.
//
//A algorithms set is:
//[
//[-1,  0, 0, 1],
//[-2, -1, 1, 2],
//[-2,  0, 0, 2]
//]

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var result [][]int
	length2 := len(nums) - 2
	length3 := length2 - 1
	end := len(nums) - 1
	for i := 0; i < length3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < length2; j ++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			left, right := j+1, end
			for left < right {
				r := nums[i] + nums[j] + nums[left] + nums[right]
				if r < target {
					left++
				} else if r > target {
					right--
				} else {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
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
	}
	return result
}
