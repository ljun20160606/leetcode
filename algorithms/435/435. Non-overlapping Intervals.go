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
func eraseOverlapIntervals(intervals []Interval) int {
	end := -1 << 31
	if len(intervals) == 1 {
		return 0
	}
	var r int

	sort.Sort(algorithms.IntervalsEnd(intervals))

	for _, v := range intervals {
		if end <= v.Start {
			end = v.End
		} else {
			r++
		}
	}

	return r
}
