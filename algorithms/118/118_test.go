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
	one int
}

type ans struct {
	one [][]int
}

func Test_Problem0118(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{0},
			ans{[][]int{}},
		},

		{
			para{1},
			ans{[][]int{
				{1},
			}},
		},

		{
			para{5},
			ans{[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, generate(p.one), "输入:%v", p)
	}
}
