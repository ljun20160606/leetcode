package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test23(t *testing.T) {
	nodes := []*algorithms.ListNode{{1, nil}, {2, nil}}
	resultOne := mergeKLists(nodes)
	expected := algorithms.Array2ListNode([]int{1, 2})
	if !reflect.DeepEqual(resultOne, expected) {
		t.Fatal("one")
	}
}
