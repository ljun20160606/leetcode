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
	one []int
}

func Test27(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{3, 2, 2, 3}, 3},
			ans{[]int{2, 2}},
		},

		{
			para{[]int{3, 1, 5, 7, 2, 2, 3}, 3},
			ans{[]int{2, 1, 5, 7, 2}},
		},

		{
			para{[]int{1, 5, 7, 2, 2}, 3},
			ans{[]int{1, 5, 7, 2, 2}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, p.one[:removeElement(p.one, p.two)], "输入:%v", p)
	}
}
