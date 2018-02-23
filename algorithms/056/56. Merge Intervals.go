package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"sort"
)

//Given a collection of Intervals, merge all overlapping Intervals.
//
//For example,
//Given [1,3],[2,6],[8,10],[15,18],
//return [1,6],[8,10],[15,18].

func merge(ins []algorithms.Interval) []algorithms.Interval {
	if len(ins) == 0 {
		return ins
	}
	sort.Sort(algorithms.Intervals(ins))
	var r []algorithms.Interval
	start, end := ins[0].Start, ins[0].End
	for _, v := range ins {
		if v.Start <= end {
			if v.End > end {
				end = v.End
			}
		} else {
			r = append(r, algorithms.Interval{Start: start, End: end})
			start = v.Start
			end = v.End
		}
	}
	r = append(r, algorithms.Interval{Start: start, End: end})
	return r
}
