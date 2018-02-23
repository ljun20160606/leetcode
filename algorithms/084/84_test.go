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
	heights []int
}

type ans struct {
	one int
}

func Test84(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				[]int{6, 5, 4, 3, 2, 1},
			},
			ans{
				12,
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5, 6},
			},
			ans{
				12,
			},
		},

		{
			para{
				[]int{2, 0, 1, 0, 1, 0},
			},
			ans{
				2,
			},
		},

		{
			para{
				[]int{2, 1, 5, 6, 2, 3},
			},
			ans{
				10,
			},
		},

		{
			para{
				[]int{4, 2, 0, 3, 2, 4, 3, 4},
			},
			ans{
				10,
			},
		},

		{
			para{
				[]int{4, 3, 4, 2, 3, 0, 2, 4},
			},
			ans{
				10,
			},
		},

		{
			para{
				[]int{},
			},
			ans{
				0,
			},
		},

		{
			para{
				[]int{1},
			},
			ans{
				1,
			},
		},

		{
			para{
				[]int{0, 0},
			},
			ans{
				0,
			},
		},

		{
			para{
				[]int{9, 0},
			},
			ans{
				9,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, largestRectangleArea(p.heights), "输入:%v", p)
		ast.Equal(a.one, naiveLargestRectangleArea(p.heights), "输入:%v", p)
	}
}
