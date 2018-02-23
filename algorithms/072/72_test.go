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
	word1 string
	word2 string
}

type ans struct {
	one int
}

func Test72(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				"esa",
				"eat",
			},
			ans{
				2,
			},
		},

		{
			para{
				"sea",
				"eat",
			},
			ans{
				2,
			},
		},

		{
			para{
				"horse",
				"ros",
			},
			ans{
				3,
			},
		},

		{
			para{
				"",
				"",
			},
			ans{
				0,
			},
		},

		{
			para{
				"gaod",
				"bad",
			},
			ans{
				2,
			},
		},

		{
			para{
				"gd",
				"bad",
			},
			ans{
				2,
			},
		},

		{
			para{
				"good",
				"bad",
			},
			ans{
				3,
			},
		},

		{
			para{
				"hello",
				"world",
			},
			ans{
				4,
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, minDistance(p.word1, p.word2), "输入:%v", p)
	}
}
