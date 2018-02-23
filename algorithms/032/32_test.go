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
	one string
}

type ans struct {
	one int
}

func Test32(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{")(()()))((((())))("},
			ans{8},
		},

		{
			para{"()(()"},
			ans{2},
		},

		{
			para{"(()())"},
			ans{6},
		},

		{
			para{"(()"},
			ans{2},
		},

		{
			para{")()())"},
			ans{4},
		},

		{
			para{"((()))"},
			ans{6},
		},

		{
			para{"((())))(((())))"},
			ans{8},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, longestValidParentheses(p.one), "输入:%v", p)
		ast.Equal(a.one, longestValidParenthesesBest(p.one), "输入:%v", p)
	}
}
