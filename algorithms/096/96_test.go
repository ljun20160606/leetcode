package algorithms

import "testing"

func Test96(t *testing.T) {
	trees := numTrees(3)
	if trees != 5 {
		t.Fatal("trees should be 5")
	}
}
