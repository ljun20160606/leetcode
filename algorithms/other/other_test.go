package other

import (
	"fmt"
	"testing"
)

func TestKsum(t *testing.T) {
	var result [][]int
	KSum([]int{-3, -1, 0, 2, 4, 5}, 1, 2, []int{}, &result)
	fmt.Println(result)
}

func TestKnapsack(t *testing.T) {
	fmt.Println(knapsack([]item{
		{
			volume: 2,
			price:  6,
		},
		{
			volume: 5,
			price:  5,
		},
		{
			volume: 3,
			price:  4,
		},
		{
			volume: 10,
			price:  15,
		},
		{
			volume: 2,
			price:  5,
		},
	}, 15))
}
