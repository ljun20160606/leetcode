# Subsets

## 描述

```txt

Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.


For example,
If nums = [1,2,3], a solution is:


[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

```

## 题解

```go
package algorithms

func subsets(nums []int) [][]int {
	var result [][]int
	helper(nums, &result, []int{}, 0)
	return result
}

func helper(nums []int, result *[][]int, temp []int, index int) {
	t := make([]int, len(temp))
	copy(t, temp)
	*result = append(*result, t)
	for i := index; i < len(nums); i++ {
		temp = append(temp, nums[i])
		helper(nums, result, temp, i+1)
		temp = temp[:len(temp)-1]
	}
}

```