package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test103(t *testing.T) {
	ast := assert.New(t)
	qst := struct {
		in       []interface{}
		expected [][]int
	}{
		in:       []interface{}{1, 2, 3, 4, nil, nil, 5},
		expected: [][]int{{1}, {3, 2}, {4, 5}},
	}

	ast.Equal(qst.expected, zigzagLevelOrder(algorithms.Array2BinaryTree(qst.in)), "输入:%v", qst.in)
}
