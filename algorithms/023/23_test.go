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
	one [][]int
}

type ans struct {
	one []int
}

func Test23(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{},
		{
			para{[][]int{
				{1, 4, 7},
			}},
			ans{[]int{1, 4, 7}},
		},
		{
			para{[][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			}},
			ans{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},

		{
			para{[][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			}},
			ans{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		},

		{
			para{[][]int{
				{1, 4, 7},
				{2, 5, 8},
				{},
			}},
			ans{[]int{1, 2, 4, 5, 7, 8}},
		},

		{
			para{[][]int{
				{1, 4, 7},
				{},
				{2, 5, 8},
			}},
			ans{[]int{1, 2, 4, 5, 7, 8}},
		},

		{
			para{[][]int{}},
			ans{[]int{}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(algorithms.Array2ListNode(a.one), mergeKLists(ArrayArray2ListNode(p.one)), "输入:%v", p)
	}
}

func ArrayArray2ListNode(nums [][]int) []*algorithms.ListNode {
	var res []*algorithms.ListNode

	for _, v := range nums {
		res = append(res, algorithms.Array2ListNode(v))
	}

	return res
}
