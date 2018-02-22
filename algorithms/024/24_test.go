package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test24(t *testing.T) {
	result := swapPairs(algorithms.Array2ListNode([]int{1, 2, 3, 4}))
	expected := algorithms.Array2ListNode([]int{2, 1, 4, 3})
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
