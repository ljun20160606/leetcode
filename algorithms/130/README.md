# Surrounded Regions

## 描述

```txt
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X


After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X


Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.

```

## 题解

```go
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

```