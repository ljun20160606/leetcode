package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test111(t *testing.T) {
	ast := assert.New(t)

	qst := []struct {
		para []interface{}
		ans  int
	}{
		{
			[]interface{}{},
			0,
		},
		{
			[]interface{}{1},
			1,
		},
		{
			[]interface{}{1, nil, 2},
			2,
		},
	}
	for _, q := range qst {
		ast.Equal(q.ans, minDepth(algorithms.Array2BinaryTree(q.para)), "输入:%v", q)
	}
}
