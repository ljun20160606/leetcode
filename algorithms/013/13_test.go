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

func Test13(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			para{"XXXIX"},
			ans{39},
		},
		{
			para{"MDCCCLXXXVIII"},
			ans{1888},
		},
		{
			para{"MCMLXXVI"},
			ans{1976},
		},
		{
			para{"MMMCMXCIX"},
			ans{3999},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, romanToInt(p.one), "输入:%v", p)
	}
}
