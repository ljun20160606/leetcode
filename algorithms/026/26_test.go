package algorithms

import (
	"reflect"
	"testing"
)

func Test26(t *testing.T) {
	nums := []int{1, 1, 2}
	newLength := removeDuplicates(nums)
	if newLength != 2 {
		t.Fail()
	}
	if !reflect.DeepEqual(nums[:newLength], []int{1, 2}) {
		t.Fail()
	}
}
