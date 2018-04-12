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
	nums []int
	k    int
}

type ans struct {
	one []int
}

func Test189(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				[]int{1, 2, 3, 4, 5, 6}, 2,
			},
			ans{
				[]int{5, 6, 1, 2, 3, 4},
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 14,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 7,
			},
			ans{
				[]int{1, 2, 3, 4, 5, 6, 7},
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5, 6, 7}, 3,
			},
			ans{
				[]int{5, 6, 7, 1, 2, 3, 4},
			},
		},

	}

	for _, q := range qs {
		a, p := q.ans, q.para
		rotate(p.nums, p.k)
		ast.Equal(a.one, p.nums, "输入:%v", p)
	}
}
