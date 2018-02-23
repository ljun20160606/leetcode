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
	two int
}

type ans struct {
	one []int
}

func Test19(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{1, 2}, 1},
			ans{[]int{1}},
		},
		{
			para{[]int{1, 2, 3, 4, 5}, 1},
			ans{[]int{1, 2, 3, 4}},
		},
		{
			para{[]int{1, 2, 3, 4, 5}, 2},
			ans{[]int{1, 2, 3, 5}},
		},
		{
			para{[]int{1}, 1},
			ans{[]int{}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(algorithms.Array2ListNode(a.one), removeNthFromEnd(algorithms.Array2ListNode(p.one), p.two), "输入:%v", p)
	}
}
