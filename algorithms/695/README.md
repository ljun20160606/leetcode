# Max Area of Island

## 描述

```txt
Given a non-empty 2D array grid of 0's and 1's, an island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.)  You may assume all four edges of the grid are surrounded by water.

Find the maximum area of an island in the given 2D array.
(If there is no island, the maximum area is 0.)

Example 1:
[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]

Given the above grid, return 6.

Note the answer is not 11, because the island must be connected 4-directionally.


Example 2:
[[0,0,0,0,0,0,0,0]]
Given the above grid, return 0.


Note:
The length of each dimension in the given grid does not exceed 50.

```

## 题解

```go
package algorithms

func maxAreaOfIsland(grid [][]int) int {
	var max int
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == 0 {
				continue
			}
			i := search(grid, x, y)
			if max < i {
				max = i
			}
		}
	}
	return max
}

func search(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	grid[x][y] = 0
	return 1 + search(grid, x-1, y) + search(grid, x+1, y) + search(grid, x, y-1) + search(grid, x, y+1)
}

```