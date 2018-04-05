package algorithms

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type question struct {
	para
	ans
}

type para struct {
	rowIndex int
}

type ans struct {
	one []int
}

func Test_Problem0119(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				3,
			},
			ans{
				[]int{1, 3, 3, 1},
			},
		},

		{
			para{
				0,
			},
			ans{
				[]int{1},
			},
		},

	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, getRow(p.rowIndex), "输入:%v", p)
	}
}
