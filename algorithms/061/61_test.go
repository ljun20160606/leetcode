package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	head []int
	k    int
}

type ans struct {
	one []int
}

func Test61(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				[]int{},
				1,
			},
			ans{
				[]int{},
			},
		},

		{
			para{
				[]int{},
				0,
			},
			ans{
				[]int{},
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5},
				2,
			},
			ans{
				[]int{4, 5, 1, 2, 3},
			},
		},

		{
			para{
				[]int{1, 2, 3, 4, 5},
				99,
			},
			ans{
				[]int{2, 3, 4, 5, 1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(algorithms.Array2ListNode(a.one), rotateRight(algorithms.Array2ListNode(p.head), p.k), "输入:%v", p)
	}
}
