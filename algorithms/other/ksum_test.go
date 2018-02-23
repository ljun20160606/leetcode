package other

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type para struct {
	nums   []int
	target int
	k      int
}

type ans struct {
	one [][]int
}

type question struct {
	p para
	a ans
}

func TestKsum(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{},
		{
			p: para{
				[]int{2, 2, 3, 4, 4},
				6,
				2,
			},
			a: ans{
				one: [][]int{{2, 4}},
			},
		},
		{
			p: para{
				[]int{2, 3, 4},
				8,
				2,
			},
		},
		{
			para{
				[]int{-3, -1, 0, 2, 4, 5},
				1,
				2,
			},
			ans{
				[][]int{{-3, 4}, {-1, 2}},
			},
		},
		{
			para{[]int{-4, -1, -1, 0, 1, 2, 2, 2, 2},
				0,
				3,
			},
			ans{[][]int{
				{-4, 2, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			}},
		},
	}

	for _, q := range qs {
		var result [][]int
		a, p := q.a, q.p
		KSum(p.nums, p.target, p.k, []int{}, &result)
		ast.Equal(a.one, result, "输入:%v", p, "输入:%v", p, "输入:%v", p)
	}
}
