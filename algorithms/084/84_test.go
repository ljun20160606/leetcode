package algorithms

import (
	"testing"
)

func Test84(t *testing.T) {
	ints := []int{2, 1, 5, 6, 2, 3}
	area := largestRectangleArea(ints)
	if area != 10 {
		t.Fail()
	}
	rectangleArea := naiveLargestRectangleArea(ints)
	if rectangleArea != 10 {
		t.Fail()
	}
}
