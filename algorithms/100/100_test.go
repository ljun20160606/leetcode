package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"testing"
)

func Test100(t *testing.T) {
	isSameOne := isSameTree(
		algorithms.Array2BinaryTree([]interface{}{1, 2, 3}),
		algorithms.Array2BinaryTree([]interface{}{1, 2, 3}),
	)

	if !isSameOne {
		t.Fatal("[1, 2, 3] and [1, 2, 3] are same")
	}

	isSameTwo := isSameTree(
		algorithms.Array2BinaryTree([]interface{}{1, 2, nil}),
		algorithms.Array2BinaryTree([]interface{}{1, nil, 3}),
	)

	if isSameTwo {
		t.Fatal("[1, 2] and [1, null, 3] are not same")
	}

	isSameThree := isSameTree(
		algorithms.Array2BinaryTree([]interface{}{1, 1, 2}),
		algorithms.Array2BinaryTree([]interface{}{1, 2, 1}),
	)

	if isSameThree {
		t.Fatal("[1, 1, 2] and [1, 2, 1] are not same")
	}
}
