package main

func gameOfLife(board [][]int) {
	ROWS, COLS := len(board), len(board[0])

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			count := neighbourCountInPlace(board, r, c)
			cur := board[r][c]

			if cur == 0 {
				if count == 3 {
					board[r][c] = 3
				}
				continue
			}

			if count < 2 || count > 3 {
				board[r][c] = 2
			}
		}
	}

	for r := 0; r < ROWS; r++ {
		for c := 0; c < COLS; c++ {
			if board[r][c] == 2 {
				board[r][c] = 0
			}
			if board[r][c] == 3 {
				board[r][c] = 1
			}
		}
	}
}

func neighbourCountInPlace(board [][]int, row, col int) int {
	count := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}

			nextRow := row + r
			if nextRow < 0 || nextRow >= len(board) {
				continue
			}

			nextCol := col + c
			if nextCol < 0 || nextCol >= len(board[0]) {
				continue
			}

			b := board[nextRow][nextCol]
			if b == 1 || b == 2 {
				count++
			}
		}
	}

	return count
}

func gameOfLifeMap(board [][]int) {
	changes := make(map[[2]int]bool)

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			count := neighbourCount(board, r, c)
			if board[r][c] == 0 {
				if count == 3 {
					changes[[2]int{r, c}] = true
				}
				continue
			}

			if count < 2 || count > 3 {
				changes[[2]int{r, c}] = false
			} else {
				changes[[2]int{r, c}] = true
			}
		}
	}

	for pos, state := range changes {
		r, c := pos[0], pos[1]
		if state {
			board[r][c] = 1
		} else {
			board[r][c] = 0
		}
	}
}

func gameOfLifeSlice(board [][]int) {
	nextState := make([][]int, len(board))
	for r := range len(board) {
		nextState[r] = make([]int, len(board[0]))
	}

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			count := neighbourCount(board, r, c)
			if board[r][c] == 0 {
				if count == 3 {
					nextState[r][c] = 1
				}
				continue
			}

			if count < 2 || count > 3 {
				nextState[r][c] = 0
			} else {
				nextState[r][c] = 1
			}
		}
	}

	copy(board, nextState)
}

func neighbourCount(board [][]int, row, col int) int {
	count := 0
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r == 0 && c == 0 {
				continue
			}

			nextRow := row + r
			if nextRow < 0 || nextRow >= len(board) {
				continue
			}

			nextCol := col + c
			if nextCol < 0 || nextCol >= len(board[0]) {
				continue
			}

			if board[nextRow][nextCol] == 1 {
				count++
			}
		}
	}

	return count
}
