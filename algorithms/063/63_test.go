package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	obstacleGrid [][]int
}

type ans struct {
	one int
}

func Test63(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				[][]int{
					{1},
				},
			},
			ans{
				0,
			},
		},

		{
			para{
				[][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			ans{
				2,
			},
		},

		{
			para{
				[][]int{
					{0, 1, 0},
					{1, 1, 0},
					{0, 0, 0},
				},
			},
			ans{
				0,
			},
		},

		{
			para{
				[][]int{},
			},
			ans{0},
		},

		{
			para{
				[][]int{
					{},
					{},
				},
			},
			ans{0},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, uniquePathsWithObstacles(p.obstacleGrid), "输入:%v", p)
	}
}
