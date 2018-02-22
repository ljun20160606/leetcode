package algorithms

import "testing"

func Test11(t *testing.T) {
	area := maxArea([]int{4, 9, 8})
	if area != 8 {
		t.Fail()
	}
}
