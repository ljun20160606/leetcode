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
	a string
	b string
}

type ans struct {
	one string
}

func Test67(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				"11",
				"11",
			},
			ans{
				"110",
			},
		},

		{
			para{
				"01",
				"000",
			},
			ans{
				"001",
			},
		},

		{
			para{
				"",
				"",
			},
			ans{
				"",
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, addBinary(p.a, p.b), "输入:%v", p)
	}
}
