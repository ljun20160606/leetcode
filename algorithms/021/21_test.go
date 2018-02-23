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
	one []int
	two []int
}

type ans struct {
	one []int
}

func Test21(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{},
				[]int{1, 3, 5, 7},
			},
			ans{[]int{1, 3, 5, 7}},
		},
		{
			para{[]int{1, 3, 5, 7},
				[]int{},
			},
			ans{[]int{1, 3, 5, 7}},
		},
		{
			para{[]int{1, 3, 5, 7},
				[]int{2, 4, 6, 8},
			},
			ans{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
		},
		{
			para{[]int{10, 20, 30},
				[]int{1, 2, 3},
			},
			ans{[]int{1, 2, 3, 10, 20, 30}},
		},
		{
			para{[]int{1, 3, 5},
				[]int{2, 4, 6, 8, 10},
			},
			ans{[]int{1, 2, 3, 4, 5, 6, 8, 10}},
		},
		{
			para{[]int{1, 3, 5, 7, 9},
				[]int{2, 4, 6},
			},
			ans{[]int{1, 2, 3, 4, 5, 6, 7, 9}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(algorithms.Array2ListNode(a.one),
			mergeTwoLists(algorithms.Array2ListNode(p.one),
				algorithms.Array2ListNode(p.two)), "输入:%v", p)
	}
}
