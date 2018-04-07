package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	prices []int
}

type ans struct {
	one int
}

func Test122(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				[]int{7, 2, 1, 5, 3, 6, 4},
			},
			ans{
				7,
			},
		},

		{
			para{
				[]int{7, 1, 5, 3, 6, 4},
			},
			ans{
				7,
			},
		},

		{
			para{
				[]int{7, 6, 5, 4, 3, 2, 1},
			},
			ans{
				0,
			},
		},

	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, maxProfit(p.prices), "输入:%v", p)
	}
}
