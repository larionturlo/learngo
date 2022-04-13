package main

//  Problem https://leetcode.com/problems/game-of-life/
func gameOfLife(board [][]int) {
	oldBoard := make([][]int, len(board))
	for i := range board {
		oldBoard[i] = make([]int, len(board[i]))
		copy(oldBoard[i], board[i])
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			alive := checkAlive(oldBoard, i, j)
			board[i][j] = alive
		}
	}
}

func checkAlive(board [][]int, i, j int) int {
	aliveNeibors := 0
	in := i - 1
	jn := j - 1
	for n := 1; n <= 9; n++ {
		if in >= 0 && jn >= 0 && in < len(board) && jn < len(board[0]) {
			if !(in == i && jn == j) {
				aliveNeibors += board[in][jn]
			}
		}

		in = i + (n%3 - 1)
		if n%3 == 0 {
			jn = j + ((9 / n) % 3)
		}
	}

	if aliveNeibors < 2 || aliveNeibors > 3 {
		return 0
	}

	if aliveNeibors == 2 {
		return board[i][j]
	}

	return 1
}
