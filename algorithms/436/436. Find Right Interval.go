package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"sort"
)

type Interval = algorithms.Interval

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
func findRightInterval(intervals []Interval) []int {
	size := len(intervals)

	starts := make([]int, size)
	idxs := make(map[int]int, size)
	res := make([]int, size)

	for i, v := range intervals {
		starts[i] = v.Start
		idxs[v.Start] = i
	}

	sort.Ints(starts)

	for i, v := range intervals {
		// SearchInts对分搜索，返回>=v.End的index
		idx := sort.SearchInts(starts, v.End)
		if idx < size {
			res[i] = idxs[starts[idx]]
		} else {
			res[i] = -1
		}
	}

	return res
}
