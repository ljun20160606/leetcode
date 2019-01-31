# Non-overlapping Intervals

## 描述

```txt

Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.


Note:

You may assume the interval's end point is always bigger than its start point.
Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.



Example 1:
Input: [ [1,2], [2,3], [3,4], [1,3] ]

Output: 1

Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.



Example 2:
Input: [ [1,2], [1,2], [1,2] ]

Output: 2

Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.



Example 3:
Input: [ [1,2], [2,3] ]

Output: 0

Explanation: You don't need to remove any of the intervals since they're already non-overlapping.


```

## 题解

```go
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
```