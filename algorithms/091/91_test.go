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
}

type ans struct {
	one int
}

func Test91(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				"27",
			},
			ans{
				1,
			},
		},
		{
			para{
				"",
			},
			ans{
				0,
			},
		},

		{
			para{
				"0",
			},
			ans{
				0,
			},
		},

		{
			para{
				"1",
			},
			ans{
				1,
			},
		},

		{
			para{
				"203",
			},
			ans{
				1,
			},
		},

		{
			para{
				"119",
			},
			ans{
				3,
			},
		},
		{
			para{
				"2011",
			},
			ans{
				2,
			},
		},

		{
			para{
				"1192",
			},
			ans{
				3,
			},
		},

		{
			para{
				"1102",
			},
			ans{
				1,
			},
		},

		{
			para{
				"1100",
			},
			ans{
				0,
			},
		},

		{
			para{
				"1190",
			},
			ans{
				0,
			},
		},

		{
			para{
				"1128",
			},
			ans{
				3,
			},
		},

		{
			para{
				"1",
			},
			ans{
				1,
			},
		},

		{
			para{
				"12",
			},
			ans{
				2,
			},
		},

		{
			para{
				"121",
			},
			ans{
				3,
			},
		},
		{
			para{
				"1212",
			},
			ans{
				5,
			},
		},

		{
			para{
				"12121",
			},
			ans{
				8,
			},
		},

		{
			para{
				"121212",
			},
			ans{
				13,
			},
		},

		{
			para{
				"1212121",
			},
			ans{
				21,
			},
		},
		{
			para{
				"12121212",
			},
			ans{
				34,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, numDecodings(p.s), "输入:%v", p)
	}
}
