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
	grid [][]int
}

type ans struct {
	one int
}

func Test64(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			}},
			ans{21},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, minPathSum(p.grid), "输入:%v", p)
	}
}
