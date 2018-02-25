package algrithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	candidates []int
	target     int
}

type ans struct {
	one [][]int
}

func Test40(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{1}, 1},
			ans{[][]int{
				{1},
			}},
		},

		{
			para{[]int{10, 1, 2, 7, 6, 1, 5}, 8},
			ans{[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, combinationSum2(p.candidates, p.target), "输入:%v", p)
	}
}
