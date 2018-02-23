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
	m int
	n int
}

type ans struct {
	one int
}

func Test62(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{},
		{
			para{
				3,
				7,
			},
			ans{
				28,
			},
		},

		{
			para{
				23,
				12,
			},
			ans{
				193536720,
			},
		},

		{
			para{
				100,
				100,
			},
			ans{
				4631081169483718960,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, uniquePaths(p.m, p.n), "输入:%v", p)
	}
}
