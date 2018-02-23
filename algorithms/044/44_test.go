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
	s string
	p string
}

type ans struct {
	one bool
}

func Test44(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				"aa",
				"a",
			},
			ans{
				false,
			},
		},

		{
			para{
				"aa",
				"aa",
			},
			ans{
				true,
			},
		},

		{
			para{
				"aaa",
				"aa",
			},
			ans{
				false,
			},
		},

		{
			para{
				"aa",
				"*",
			},
			ans{
				true,
			},
		},

		{
			para{
				"aa",
				"a*",
			},
			ans{
				true,
			},
		},

		{
			para{
				"ab",
				"?*",
			},
			ans{
				true,
			},
		},

		{
			para{
				"aab",
				"c*a*b",
			},
			ans{
				false,
			},
		},

		{
			para{
				"ab",
				"*",
			},
			ans{
				true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, wildCardIsMatch(p.s, p.p), "输入:%v", p)
	}
}
