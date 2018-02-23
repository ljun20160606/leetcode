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
}

type ans struct {
	one [][]int
}

func Test15(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{1, -1, -1, 0}},
			ans{[][]int{
				{-1, 0, 1},
			}},
		},

		{
			para{[]int{-1, 0, 1, 2, 2, 2, 2, -1, -4}},
			ans{[][]int{
				{-4, 2, 2},
				{-1, -1, 2},
				{-1, 0, 1},
			}},
		},
		{
			para{[]int{0, 0, 0, 0, 0}},
			ans{[][]int{
				{0, 0, 0},
			}},
		},
		{
			para{[]int{1, 1, -2}},
			ans{[][]int{
				{-2, 1, 1},
			}},
		},
		{
			para{[]int{0, 0, 0}},
			ans{[][]int{
				{0, 0, 0},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, threeSum(p.one), "输入:%v", p)
	}
}
