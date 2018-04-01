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
	m   int
	two []int
	n   int
}

type ans struct {
	one []int
}

func Test_Problem0088(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{0}, 0, []int{1}, 1},
			ans{[]int{1}},
		},

		{
			para{[]int{1, 0}, 1, []int{2}, 1},
			ans{[]int{1, 2}},
		},

		{
			para{[]int{1, 3, 5, 7, 0, 0}, 4, []int{2, 2}, 2},
			ans{[]int{1, 2, 2, 3, 5, 7}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		merge(p.one, p.m, p.two, p.n)
		ast.Equal(a.one, p.one, "输入:%v", p)
	}
}
