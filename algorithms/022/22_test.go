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
	one int
}

type ans struct {
	one []string
}

func Test22(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{2},
			ans{[]string{
				"(())",
				"()()",
			}},
		},

		{
			para{3},
			ans{[]string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			}},
		},

		{
			para{4},
			ans{[]string{
				"(((())))",
				"((()()))",
				"((())())",
				"((()))()",
				"(()(()))",
				"(()()())",
				"(()())()",
				"(())(())",
				"(())()()",
				"()((()))",
				"()(()())",
				"()(())()",
				"()()(())",
				"()()()()",
			}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, generateParenthesis(p.one), "输入:%v", p)
	}
}
