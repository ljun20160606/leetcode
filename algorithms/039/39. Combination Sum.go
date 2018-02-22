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
