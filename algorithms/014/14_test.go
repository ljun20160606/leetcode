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
	one []string
}

type ans struct {
	one string
}

func Test14(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{
				[]string{"abcdd", "abcde", "ab"},
			},
			ans{"ab"},
		},
		{
			para{
				[]string{"abcdd", "abcde"},
			},
			ans{"abcd"},
		},
		{
			para{
				[]string{},
			},
			ans{""},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, longestCommonPrefix(p.one), "输入:%v", p)
	}
}
