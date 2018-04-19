package algorithms

import "github.com/ljun20160606/leetcode/algorithms"

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval = algorithms.Interval

func insert(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	ret := make([]Interval, 0)
	prev := newInterval
	idx := 0
	for idx < len(intervals) {
		if prev.Start > intervals[idx].Start {
			if isIntersect(intervals[idx], prev) {
				prev = merge(intervals[idx], prev)
			} else {
				ret = append(ret, intervals[idx])
			}
		} else {
			if isIntersect(prev, intervals[idx]) {
				prev = merge(prev, intervals[idx])
			} else {
				ret = append(ret, prev)
				prev = intervals[idx]
			}
		}
		idx++
	}
	ret = append(ret, prev)
	return ret
}

func isIntersect(a Interval, b Interval) bool {
	return a.End >= b.Start
}

func merge(a Interval, b Interval) Interval {
	if a.End < b.End {
		return Interval{Start: a.Start, End: b.End}
	}
	return a
}
