package algorithms

import (
	"fmt"
	"testing"
)

func Test31(t *testing.T) {
	i := []int{1, 3, 2}
	nextPermutation(i)
	fmt.Println(i)
}
