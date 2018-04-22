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
