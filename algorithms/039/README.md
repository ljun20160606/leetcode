# Combination Sum

## 描述

Given a set of candidate numbers (C) (without duplicates) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.

The same repeated number may be chosen from C unlimited number of times.

Note:

* All numbers (including target) will be positive integers.
* The solution set must not contain duplicate combinations.
For example, given candidate set [2, 3, 6, 7] and target 7, 
A solution set is:

```json 
[
  [7],
  [2, 2, 3]
]
```

## 题解

```go
package algrithms

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var helper func(candidate, share []int, target int)
	var r [][]int
	helper = func(c, share []int, target int) {
		if target == 0 {
			r = append(r, share)
			return
		}
		if len(c) == 0 || target < c[0] {
			return
		}
		share = share[:len(share):len(share)]

		helper(c, append(share, c[0]), target-c[0])
		helper(c[1:], share, target)
	}
	helper(candidates, []int{}, target)
	return r
}
```