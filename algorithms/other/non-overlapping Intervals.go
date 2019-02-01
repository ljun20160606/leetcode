package other

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"sort"
)

type Interval = algorithms.Interval

// 最大调度区间问题
// 存在多个时间区间{start, end}，求最大工作区间数
// 贪心、根据end时间排序
func numbersOfNonOverlapping(intervals []Interval) int {
	sort.Sort(algorithms.IntervalsEnd(intervals))
	var pre int
	var r int
	for i := range intervals {
		interval := intervals[i]
		if pre < interval.Start {
			pre = interval.End
			r++
		}
	}
	return r
}