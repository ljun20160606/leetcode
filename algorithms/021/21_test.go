package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test21(t *testing.T) {
	result := mergeTwoLists(algorithms.Array2ListNode([]int{2, 4}), algorithms.Array2ListNode([]int{1, 3}))
	expected := algorithms.Array2ListNode([]int{1, 2, 3, 4})
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
