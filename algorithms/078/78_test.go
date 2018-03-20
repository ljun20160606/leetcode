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
	nums []int
}

type ans struct {
	one [][]int
}

func Test78(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{
			para{
				[]int{1, 2, 3},
			},
			ans{[][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, subsets(p.nums), "输入:%v", p)
	}
}
