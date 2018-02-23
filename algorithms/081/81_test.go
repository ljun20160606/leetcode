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
	nums   []int
	target int
}

type ans struct {
	one bool
}

func Test81(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				[]int{4, 4, 5, 5, 6, 6, 7, 7, 0, 0, 1, 1, 2, 2},
				6,
			},
			ans{
				true,
			},
		},

		{
			para{
				[]int{4, 5, 6, 7, 0, 1, 2},
				6,
			},
			ans{
				true,
			},
		},

		{
			para{
				[]int{4, 5, 6, 7, 0, 1, 2},
				3,
			},
			ans{
				false,
			},
		},

		{
			para{
				[]int{6, 7, 0, 1, 3, 3, 3},
				3,
			},
			ans{
				true,
			},
		},

		{
			para{
				[]int{6, 7, 0, 1, 3, 3, 3},
				7,
			},
			ans{
				true,
			},
		},

		{
			para{
				[]int{2, 2, 2, 1, 1, 1, 1},
				7,
			},
			ans{
				false,
			},
		},

		{
			para{
				[]int{},
				3,
			},
			ans{
				false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, search(p.nums, p.target), "输入:%v", p)
	}
}
