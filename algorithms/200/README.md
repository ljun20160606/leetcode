# Number of Islands

## 描述

```txt
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:
11110110101100000000
Answer: 1
Example 2:
11000110000010000011
Answer: 3

Credits:Special thanks to @mithmatt for adding this problem and creating all test cases.
```

## 题解

```go
package algorithms

func numIslands(grid [][]byte) int {
	var count int
	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '1' {
				search(grid, x, y)
				count++
			}
		}
	}
	return count
}

func search(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == '0' {
		return
	}
	grid[x][y] = '0'
	search(grid, x-1, y)
	search(grid, x+1, y)
	search(grid, x, y-1)
	search(grid, x, y+1)
}

```