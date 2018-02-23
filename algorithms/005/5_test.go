package algorithms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSym(t *testing.T) {
	assert.False(t, sym("abbbc"))
}

type para struct {
	one string
}

type ans struct {
	one string
}

type question struct {
	p para
	a ans
}

func Test5(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para{
				one: "",
			},
			a: ans{
				one: "",
			},
		},
		{
			p: para{
				one: "babad",
			},
			a: ans{
				one: "bab",
			},
		},
		{
			p: para{
				one: "cbbd",
			},
			a: ans{
				one: "bb",
			},
		},
		{
			p: para{
				one: "abbcccddcccbba",
			},
			a: ans{
				one: "abbcccddcccbba",
			},
		},
		{
			p: para{
				one: "a",
			},
			a: ans{
				one: "a",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, longestPalindrome(p.one), "输入:%v", p)
	}
}
