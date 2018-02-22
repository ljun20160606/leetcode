package algorithms

import (
	"math"
	"testing"
)

func Test29(t *testing.T) {
	if divide(6, 0) != math.MaxInt32 {
		t.Fail()
	}
	if divide(6, 2) != 3 {
		t.Fail()
	}
}
