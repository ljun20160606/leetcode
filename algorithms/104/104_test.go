package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test104(t *testing.T) {
	ast := assert.New(t)

	tcs := []struct {
		pre, in []int
		ans     int
	}{
		{
			[]int{},
			[]int{},
			0,
		},

		{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			3,
		},

		{
			[]int{3, 9, 20, 15, 10, 7},
			[]int{9, 3, 10, 15, 20, 7},
			4,
		},
	}

	for _, tc := range tcs {
		ast.Equal(tc.ans, maxDepth(algorithms.PreIn2Tree(tc.pre, tc.in)), "输入:%v", tc)
	}
}
