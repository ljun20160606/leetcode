package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"testing"
)

func Test101(t *testing.T) {
	isSymmetricOne := isSymmetric(algorithms.Array2BinaryTree([]interface{}{1, 2, 2, 3, 4, 4, 3}))
	if !isSymmetricOne {
		t.Fail()
	}
	isSymmetricTwo := isSymmetric(nil)
	if !isSymmetricTwo {
		t.Fail()
	}
	isSymmetricThree := isSymmetric(algorithms.Array2BinaryTree([]interface{}{1, 2, 3}))
	if isSymmetricThree {
		t.Fail()
	}
}
