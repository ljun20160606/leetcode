package algorithms

//Follow up for "Unique Paths":
//
//Now consider if some obstacles are added to the grids. How many unique paths would there be?
//
//An obstacle and empty space is marked as 1 and 0 respectively in the grid.
//
//For example,
//There is one obstacle in the middle of a 3x3 grid as illustrated below.
//
//[
//	[0,0,0],
//	[0,1,0],
//	[0,0,0]
//]
//The total number of unique paths is 2.
//
//Note: m and n will be at most 100.

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	lenX, lenY := len(obstacleGrid[0]), len(obstacleGrid)
	if obstacleGrid[lenY-1][lenX-1] == 1 {
		return 0
	}
	matrix := make([][]int, lenY)
	flag := true
	for y := range matrix {
		matrix[y] = make([]int, lenX)
		if obstacleGrid[y][0] == 1 {
			flag = false
		}
		if flag {
			matrix[y][0] = 1
		}
	}
	flag = true
	for x := range matrix[0] {
		if obstacleGrid[0][x] == 1 {
			flag = false
		}
		if flag {
			matrix[0][x] = 1
		}
	}
	for y := 1; y < lenY; y++ {
		for x := 1; x < lenX; x++ {
			if obstacleGrid[y][x] != 1 {
				matrix[y][x] = matrix[y][x-1] + matrix[y-1][x]
			}
		}
	}
	return matrix[lenY-1][lenX-1]
}
