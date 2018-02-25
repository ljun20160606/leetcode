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
