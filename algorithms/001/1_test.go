package algorithms

import (
	"reflect"
	"testing"
)

func Test1(t *testing.T) {
	if !reflect.DeepEqual(twoSum([]int{1, 3, 2}, 3), []int{0, 2}) {
		t.Fatal("result must be 0, 2")
	}
	if len(twoSum([]int{}, 1)) != 0 {
		t.Fatal("result must be []")
	}
	if len(twoSum([]int{1, 3, 2}, 6)) != 0 {
		t.Fatal("result must be []")
	}
}

func Test1V1(t *testing.T) {
	if !reflect.DeepEqual(twoSumV0([]int{1, 2, 3, 5, 6}, 5), []int{1, 2}) {
		t.Fatal("result must be 1, 2")
	}
	if len(twoSumV0([]int{}, 1)) != 0 {
		t.Fatal("result must be []")
	}
	if len(twoSumV0([]int{1, 3, 2}, 6)) != 0 {
		t.Fatal("result must be []")
	}
}
