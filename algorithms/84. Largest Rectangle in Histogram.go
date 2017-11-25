package algorithms

import "github.com/LFZJun/leetcode/algorithms/util"

//Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.
//
//
//Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
//
//
//The largest rectangle is shown in the shaded area, which has area = 10 unit.
//
//For example,
//Given heights = [2,1,5,6,2,3],
//return 10.

func naiveLargestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	size := len(heights)
	left, right := make([]int, size), make([]int, size)

	for i := 1; i < size; i++ {
		for j := i - 1; j >= 0 && heights[j] >= heights[i]; j-- {
			left[i]++
		}
	}

	for i := size - 2; i >= 0; i-- {
		for j := i + 1; j < size && heights[j] >= heights[i]; j++ {
			right[i]++
		}
	}

	var maxSquare int
	for i := 0; i < size; i++ {
		maxSquare = util.MaxOfTwo(maxSquare, (left[i]+right[i]+1)*heights[i])
	}
	return maxSquare
}
