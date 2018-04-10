package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_canFinish(t *testing.T) {
	ast := assert.New(t)

	var tcs = []struct {
		numCourses    int
		prerequisites [][]int
		ans           bool
	}{
		{
			4,
			[][]int{
				{1, 0},
				{2, 1},
				{2, 0},
				{3, 0},
				{3, 1},
				{3, 2},
			},
			true,
		},

		{
			4,
			[][]int{
				{1, 0},
				{2, 1},
				{2, 0},
				{3, 0},
				{3, 1},
				{3, 2},
				{2, 3},
			},
			false,
		},

		{
			2,
			[][]int{
				{0, 1},
			},
			true,
		},

		{
			2,
			[][]int{
				{1, 0},
			},
			true,
		},

		{
			2,
			[][]int{
				{1, 0},
				{0, 1},
			},
			false,
		},
	}

	for _, tc := range tcs {
		ast.Equal(tc.ans, canFinish(tc.numCourses, tc.prerequisites), "输入:%v", tc)
	}
}
