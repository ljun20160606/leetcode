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
	two string
}

type ans struct {
	one int
}

func Test28(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{"", ""},
			ans{0},
		},

		{
			para{"abcd", "bc"},
			ans{1},
		},

		{
			para{"abcde", "c"},
			ans{2},
		},

		{
			para{"abcde", "f"},
			ans{-1},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, strStr(p.one, p.two), "输入:%v", p)
	}
}
