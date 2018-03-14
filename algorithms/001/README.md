# Two Sum

## 描述

```txt
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.


Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] &#43; nums[1] = 2 &#43; 7 = 9,
return [0, 1].


```

## 题解

```go
package algorithms

//Given nums = [2, 7, 11, 15], target = 9,
//
//Because nums[0] + nums[1] = 2 + 7 = 9,
//return [0, 1].

func twoSum(nums []int, target int) []int {
	t := make(map[int]int, len(nums))
	for i, v := range nums {
		if vv, ok := t[v]; ok {
			return []int{vv, i}
		}
		t[target-v] = i
	}
	return nil
}

```