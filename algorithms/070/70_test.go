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

func Test70(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				0,
			},
			ans{
				0,
			},
		},

		{
			para{
				10,
			},
			ans{
				89,
			},
		},

		{
			para{
				44,
			},
			ans{
				1134903170,
			},
		},

		{
			para{
				1,
			},
			ans{
				1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, climbStairs(p.n), "输入:%v", p)
	}
}
