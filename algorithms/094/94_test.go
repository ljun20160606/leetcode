package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test94(t *testing.T) {
	r := inorderTraversal(algorithms.Array2BinaryTree([]interface{}{1, nil, 2, nil, nil, 3}))
	if !reflect.DeepEqual(r, []int{1, 3, 2}) {
		t.Fatal(r)
	}
}
