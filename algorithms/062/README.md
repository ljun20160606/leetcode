# Unique Paths

## 描述

```txt
A robot is located at the top-left corner of a m x n grid (marked &#39;Start&#39; in the diagram below).

The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked &#39;Finish&#39; in the diagram below).

How many possible unique paths are there?



Above is a 3 x 7 grid. How many possible unique paths are there?


Note: m and n will be at most 100.
```

## 题解

```go
package algorithms

//A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
//
//The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).
//
//How many possible unique paths are there?

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	matrix := make([][]int, n)
	for y := range matrix {
		matrix[y] = make([]int, m)
		matrix[y][0] = 1
	}
	for x := range matrix[0] {
		matrix[0][x] = 1
	}
	for y := 1; y < n; y++ {
		for x := 1; x < m; x++ {
			matrix[y][x] = matrix[y][x-1] + matrix[y-1][x]
		}
	}
	return matrix[n-1][m-1]
}

```