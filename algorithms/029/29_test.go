package algorithms

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	one int
	two int
}

type ans struct {
	one int
}

func Test29(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{2, 3},
			ans{0},
		},

		{
			para{2, 0},
			ans{math.MaxInt32},
		},

		{
			para{1024, 3},
			ans{341},
		},

		{
			para{1024, -3},
			ans{-341},
		},

		{
			para{-1024, 3},
			ans{-341},
		},

		{
			para{-1024, -3},
			ans{341},
		},

		{
			para{1024, 1},
			ans{1024},
		},

		{
			para{math.MaxInt32, -1},
			ans{math.MinInt32},
		},

		{
			para{2147483648, 1},
			ans{math.MaxInt32},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, divide(p.one, p.two), "输入:%v", p)
	}
}
