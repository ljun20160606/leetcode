package algorithms

import (
	"github.com/ljun20160606/leetcode/algorithms"
	"reflect"
	"testing"
)

func Test206(t *testing.T) {
	if !reflect.DeepEqual(
		reverseList(algorithms.Array2ListNode([]int{1, 2})),
		algorithms.Array2ListNode([]int{2, 1}),
	) {
		t.Fatal("not equal")
	}
}
