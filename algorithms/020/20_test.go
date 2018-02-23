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
	one bool
}

func Test20(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{""},
			ans{true},
		},
		{
			para{")"},
			ans{false},
		},
		{
			para{"}"},
			ans{false},
		},
		{
			para{"()[]{}"},
			ans{true},
		},
		{
			para{"(]"},
			ans{false},
		},
		{
			para{"({[]})"},
			ans{true},
		},
		{
			para{"(){[({[]})]}"},
			ans{true},
		},
		{
			para{"((([[[{{{"},
			ans{false},
		},
		{
			para{"(())]]"},
			ans{false},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, isValid(p.one), "输入:%v", p)
	}
}
