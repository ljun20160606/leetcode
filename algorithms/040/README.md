# Combination Sum II

## 描述

    Given a collection of candidate numbers (C) and a target number (T), find all unique combinations in C where the candidate numbers sums to T.
    
    Each number in C may only be used once in the combination.
    
    Note:
    All numbers (including target) will be positive integers.
    The solution set must not contain duplicate combinations.
    For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8, 
    A solution set is: 
    [
      [1, 7],
      [1, 2, 5],
      [2, 6],
      [1, 1, 6]
    ]

## 题解

```go
package algrithms

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var r [][]int
	var helper func(c, share []int, t int)
	helper = func(newCandidates, share []int, newTarget int) {
		if newTarget == 0 {
			r = append(r, share)
			return
		}
		if len(newCandidates) == 0 || newTarget < newCandidates[0] {
			return
		}
		share = share[:len(share):len(share)]
		helper(newCandidates[1:], append(share, newCandidates[0]), newTarget-newCandidates[0])
		var i int
		for k := range newCandidates {
			if newCandidates[k] != newCandidates[0] {
				i = k
				break
			} else {
				i++
			}
		}
		newCandidates = newCandidates[i:]
		helper(newCandidates, share, newTarget)
	}
	helper(candidates, []int{}, target)
	return r
}

```