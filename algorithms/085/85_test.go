package algorithms

import (
	"testing"
)

func Test85(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	naiveRectangle := naiveMaximalRectangle(matrix)
	if naiveRectangle != 6 {
		t.Fail()
	}
}
