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

func Test16(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{[]int{-1, 2, 1, -4}, 1},
			ans{2},
		},
		{
			para{[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, 1},
			ans{0},
		},
		{
			para{[]int{-1, 2, 2, 2, 2, 2, 2, 2, 1, -4}, 0},
			ans{0},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, threeSumClosest(p.one, p.two), "输入:%v", p)
	}
}
