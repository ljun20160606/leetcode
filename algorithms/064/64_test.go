package algorithms

import (
	"fmt"
	"testing"
)

func Test64(t *testing.T) {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}
