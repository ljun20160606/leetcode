package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test25(t *testing.T) {
	result := reverseKGroup(algorithms.Array2ListNode([]int{1, 2, 3, 4, 5}), 2)
	expected := algorithms.Array2ListNode([]int{2, 1, 4, 3, 5})
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}
