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
	one int
}

func Test53(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			ans{6},
		},

		{
			para{[]int{-2}},
			ans{-2},
		},

		{
			para{[]int{}},
			ans{0},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, maxSubArray(p.one), "输入:%v", p)
	}
}
