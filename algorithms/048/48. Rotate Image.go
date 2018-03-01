package algrithms

func rotate(matrix [][]int) {
	for i, temp, n := 0, 0, len(matrix)-1; i <= n/2; i++ {
		for j := i; j < n-i; j++ {
			temp = matrix[j][n-i]
			matrix[j][n-i] = matrix[i][j]
			matrix[i][j] = matrix[n-j][i]
			matrix[n-j][i] = matrix[n-i][n-j]
			matrix[n-i][n-j] = temp
		}
	}
}
