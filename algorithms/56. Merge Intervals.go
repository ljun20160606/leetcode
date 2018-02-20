package algorithms

import (
	"sort"
)

//Given a collection of intervals, merge all overlapping intervals.
//
//For example,
//Given [1,3],[2,6],[8,10],[15,18],
//return [1,6],[8,10],[15,18].

type interval struct {
	Start int
	End   int
}

type intervals []interval

func (in intervals) Len() int {
	return len(in)
}

func (in intervals) Less(i, j int) bool {
	x, y := in[i], in[j]
	return x.Start < y.Start
}

func (in intervals) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func merge(ins []interval) []interval {
	if len(ins) == 0 {
		return ins
	}
	sort.Sort(intervals(ins))
	var r []interval
	start, end := ins[0].Start, ins[0].End
	for _, v := range ins {
		if v.Start <= end {
			if v.End > end {
				end = v.End
			}
		} else {
			r = append(r, interval{start, end})
			start = v.Start
			end = v.End
		}
	}
	r = append(r, interval{start, end})
	return r
}
