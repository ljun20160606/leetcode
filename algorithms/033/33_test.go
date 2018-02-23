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
	one int
}

func Test33(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0, 1, 2}, 5},
			ans{1},
		},

		{
			para{[]int{6, 7, 0, 1, 2, 3, 4, 5}, 1},
			ans{3},
		},

		{
			para{[]int{6, 7, 0, 1, 2, 3}, 3},
			ans{5},
		},

		{
			para{[]int{6, 7, 0, 1, 2, 3}, 6},
			ans{0},
		},

		{
			para{[]int{1, 3}, 0},
			ans{-1},
		},

		{
			para{[]int{}, 5},
			ans{-1},
		},

		{
			para{[]int{2}, 5},
			ans{-1},
		},

		{
			para{[]int{2}, 2},
			ans{0},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, search(p.one, p.two), "输入:%v", p)
	}
}
