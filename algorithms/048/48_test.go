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
	one [][]int
}

type ans struct {
	one [][]int
}

func Test48(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[][]int{
				{1, 2},
				{3, 4},
			}},
			ans{[][]int{
				{3, 1},
				{4, 2},
			}},
		},

		{
			para{[][]int{
				{1, 1, 1, 1},
				{2, 2, 2, 2},
				{3, 3, 3, 3},
				{4, 4, 4, 4},
			}},
			ans{[][]int{
				{4, 3, 2, 1},
				{4, 3, 2, 1},
				{4, 3, 2, 1},
				{4, 3, 2, 1},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para
		rotate(p.one)
		ast.Equal(a.one, p.one, "输入:%v", p)
	}
}
