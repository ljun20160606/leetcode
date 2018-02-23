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
	n int
}

type ans struct {
	one int
}

func Test96(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				1,
			},
			ans{
				1,
			},
		},

		{
			para{
				2,
			},
			ans{
				2,
			},
		},

		{
			para{
				3,
			},
			ans{
				5,
			},
		},

		{
			para{
				4,
			},
			ans{
				14,
			},
		},

		{
			para{
				5,
			},
			ans{
				42,
			},
		},

		{
			para{
				6,
			},
			ans{
				132,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, numTrees(p.n), "输入:%v", p, "输入:%v", p)
	}
}
