# 3Sum Closest

## 描述

```txt
Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

    For example, given array S = {-1 2 1 -4}, and target = 1.

    The sum that is closest to the target is 2. (-1 &#43; 2 &#43; 1 = 2).

```

## 题解

```go
package algorithms

import (
	"sort"
)

//Given an array S of n integers, find three integers in S such that the sum is closest to a given number, target. Return the sum of the three integers. You may assume that each input would have exactly one algorithms.
//
//For example, given array S = {-1 2 1 -4}, and target = 1.
//
//The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	length := len(nums) - 2
	end := len(nums) - 1
	var closest, distance, d int
	for i := 0; i < length; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, end
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < target {
				d = target - sum
				left++
			} else if sum > target {
				d = sum - target
				right--
			} else {
				return sum
			}
			if d < distance || distance == 0 {
				closest = sum
				distance = d
			}
		}
	}
	return closest
}

```