package algorithms

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		search(board, i, 0)
		search(board, i, n-1)
	}
	for i := 0; i < n; i++ {
		search(board, 0, i)
		search(board, m-1, i)
	}
	for x := range board {
		for y := range board[x] {
			b := &board[x][y]
			switch *b {
			case '#':
				*b = 'O'
			case 'O':
				*b = 'X'
			}
		}
	}
}

func search(board [][]byte, x, y int) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != 'O' {
		return
	}
	board[x][y] = '#'
	search(board, x-1, y)
	search(board, x+1, y)
	search(board, x, y-1)
	search(board, x, y+1)
}
