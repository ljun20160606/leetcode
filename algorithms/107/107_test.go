package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test107(t *testing.T) {
	result := levelOrderBottom(algorithms.Array2BinaryTree([]interface{}{3, 9, 20, nil, nil, 15, 7}))
	expected := [][]int{{15, 7}, {9, 20}, {3}}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("need %v but get %v", expected, result)
	}
	if levelOrderBottom(nil) != nil {
		t.Fatal("must be nil")
	}
}
