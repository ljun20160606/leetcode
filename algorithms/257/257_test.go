package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test257(t *testing.T) {
	ast := assert.New(t)

	qst := []struct {
		para []interface{}
		ans  []string
	}{
		{
			[]interface{}{},
			nil,
		},
		{
			[]interface{}{1, 2, 3, nil, 5},
			[]string{"1->2->5", "1->3"},
		},
	}
	for _, q := range qst {
		ast.Equal(q.ans, binaryTreePaths(algorithms.Array2BinaryTree(q.para)), "输入:%v", q)
	}
}
