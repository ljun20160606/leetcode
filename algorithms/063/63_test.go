package algorithms

import (
	"fmt"
	"testing"
)

func Test63(t *testing.T) {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}
