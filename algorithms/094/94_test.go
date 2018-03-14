package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}

type para struct {
	pre []interface{}
}

type ans struct {
	one []int
}

func Test94(t *testing.T) {
	ast := assert.New(t)

	qs := []question{

		{},

		{
			para{
				[]interface{}{1, 2, 3},
			},
			ans{
				[]int{2, 1, 3},
			},
		},
	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, inorderTraversal(algorithms.Array2BinaryTree(p.pre)), "输入:%v", p)
	}
}
