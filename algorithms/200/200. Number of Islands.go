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
