package algorithms

// 1 0 0 1
// 0 1 1 0
// 0 1 1 1
// 1 0 1 1
func findCircleNum(M [][]int) int {
	var count int
	visited := make([]bool, len(M))
	for i := range M {
		if visited[i] {
			continue
		}
		search(M, i, &visited)
		count++
	}
	return count
}

func search(M [][]int, i int, visitedPtr *[]bool) {
	visited := *visitedPtr
	visited[i] = true
	// Cause M is a square
	for j := range M {
		if M[i][j] == 0 || visited[j] {
			continue
		}
		search(M, j, visitedPtr)
	}
}
