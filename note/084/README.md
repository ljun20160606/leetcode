# Largest Rectangle in Histogram

## 描述

> Given n non-negative integers representing the histogram's bar height where the width of each bar is 1, find the area of largest rectangle in the histogram.

![histogram](histogram.png)

> Above is a histogram where width of each bar is 1, given height = `[2,1,5,6,2,3]`.

![histogram_area](histogram_area.png)

> The largest rectangle is shown in the shaded area, which has area = 10 unit.
>
> For example, Given heights = `[2,1,5,6,2,3]`, return 10.

## 题解

### 1. 朴素做法，左右扫描

> 遍历`heights[i]`, 向左遍历获取最左边界存入`left[i]`, 向右扫描获取最右边界存入`right[i]`, 遍历`heights[i]`, `s=(left[i]+right[i]+1)*heights[i]`, 遍历的同时替换最大面积。

````go
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
````

### 2. 递增栈

````go
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	heights = append(heights, 0)
	var square int
	for i, stack, top := 0, make([]int, len(heights)), -1; i < len(heights); i++ {
		for top != -1 && heights[i] <= heights[stack[top]] {
			height := heights[stack[top]]
			top--
			var width int
			if top == -1 {
				width = i
			} else {
				width = i - 1 - stack[top]
			}
			square = util.MaxOfTwo(square, height*width)
		}
		top++
		stack[top] = i
	}
	return square
}
````