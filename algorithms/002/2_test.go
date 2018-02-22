package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test2(t *testing.T) {
	result := addTwoNumbers(
		algorithms.Array2ListNode([]int{2, 4, 3}),
		algorithms.Array2ListNode([]int{5, 6, 4}),
	)
	expected := algorithms.Array2ListNode([]int{7, 0, 8})
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
