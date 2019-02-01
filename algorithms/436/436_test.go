package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tcs = []struct {
	intervals []Interval
	ans       []int
}{

	{
		algorithms.Int2sToIntervals([][]int{
			{3, 4},
			{2, 3},
			{1, 2},
		}),
		[]int{-1, 0, 1},
	},

	{
		algorithms.Int2sToIntervals([][]int{
			{1, 2},
		}),
		[]int{-1},
	},

	{
		algorithms.Int2sToIntervals([][]int{
			{1, 4},
			{2, 3},
			{3, 4},
		}),
		[]int{-1, 2, -1},
	},
}

func Test436(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, findRightInterval(tc.intervals), "输入:%v", tc)
	}
}
