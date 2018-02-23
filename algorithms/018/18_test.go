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
	one []int
	two int
}

type ans struct {
	one [][]int
}

func Test18(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{1, 0, -1, 0, -2, 2}, 0},
			ans{[][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			}},
		},
		{
			para{[]int{0, 0, 0, 0}, 0},
			ans{[][]int{
				{0, 0, 0, 0},
			}},
		},
		{
			para{[]int{-1, -1, -2, -2, 1, 1, 2, 2, 0, 0, 0, 0}, 0},
			ans{[][]int{
				{-2, -2, 2, 2},
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-2, 0, 1, 1},
				{-1, -1, 0, 2},
				{-1, -1, 1, 1},
				{-1, 0, 0, 1},
				{0, 0, 0, 0},
			}},
		},
		{
			para{[]int{0, -1, 0, 1, -2, -5, 3, 5, 0}, 6},
			ans{[][]int{
				{-2, 0, 3, 5},
				{0, 0, 1, 5},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, fourSum(p.one, p.two), "输入:%v", p)
	}
}
