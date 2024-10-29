package main

import (
	"slices"
)

func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]struct{}, 0, 9)
	cols := make([]map[byte]struct{}, 0, 9)
	grids := make([]map[byte]struct{}, 0, 9)

	for range 9 {
		rows = append(rows, make(map[byte]struct{}))
		cols = append(cols, make(map[byte]struct{}))
		grids = append(grids, make(map[byte]struct{}))
	}

	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}

			gridIdx := (r/3)*2 + (r/3 + c/3)

			if _, ok := rows[r][board[r][c]]; ok {
				return false
			}

			if _, ok := cols[c][board[r][c]]; ok {
				return false
			}

			if _, ok := grids[gridIdx][board[r][c]]; ok {
				return false
			}

			rows[r][board[r][c]] = struct{}{}
			cols[c][board[r][c]] = struct{}{}
			grids[gridIdx][board[r][c]] = struct{}{}
		}
	}

	return true
}

func isValidSudokuImproved(board [][]byte) bool {
	numMap := make(map[[2]int]byte)

	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}
			numMap[[2]int{r, c}] = board[r][c]
		}
	}

	for r := range 9 {
		row := make([]byte, 0)
		col := make([]byte, 0)
		for c := range 9 {
			if v, ok := numMap[[2]int{r, c}]; ok {
				row = append(row, v)
			}

			if v, ok := numMap[[2]int{c, r}]; ok {
				col = append(col, v)
			}
		}

		slices.SortFunc(row, func(a, b byte) int {
			return int(a) - int(b)
		})

		slices.SortFunc(col, func(a, b byte) int {
			return int(a) - int(b)
		})

		for i := 0; i < len(row)-1; i++ {
			if row[i] == row[i+1] {
				return false
			}
		}

		for i := 0; i < len(col)-1; i++ {
			if col[i] == col[i+1] {
				return false
			}
		}
	}

	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			grid := make([]byte, 0)
			for k := range 3 {
				for j := range 3 {
					if v, ok := numMap[[2]int{r + k, c + j}]; ok {
						grid = append(grid, v)
					}
				}
			}

			slices.SortFunc(grid, func(a, b byte) int {
				return int(a) - int(b)
			})

			for i := 0; i < len(grid)-1; i++ {
				if grid[i] == grid[i+1] {
					return false
				}
			}
		}

	}

	return true
}

func isValidSudokuFirstApproach(board [][]byte) bool {
	numMap := make(map[[2]int]byte)

	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}
			numMap[[2]int{r, c}] = board[r][c]
		}
	}

	for r := range 9 {
		row := make([]byte, 0)
		for c := range 9 {
			if v, ok := numMap[[2]int{r, c}]; ok {
				row = append(row, v)
			}
		}

		slices.SortFunc(row, func(a, b byte) int {
			return int(a) - int(b)
		})

		for i := 0; i < len(row)-1; i++ {
			if row[i] == row[i+1] {
				return false
			}
		}
	}

	for c := range 9 {
		col := make([]byte, 0)
		for r := range 9 {
			if v, ok := numMap[[2]int{r, c}]; ok {
				col = append(col, v)
			}
		}

		slices.SortFunc(col, func(a, b byte) int {
			return int(a) - int(b)
		})

		for i := 0; i < len(col)-1; i++ {
			if col[i] == col[i+1] {
				return false
			}
		}
	}

	for r := 0; r < 9; r += 3 {
		for c := 0; c < 9; c += 3 {
			grid := make([]byte, 0)
			for k := range 3 {
				for j := range 3 {
					if v, ok := numMap[[2]int{r + k, c + j}]; ok {
						grid = append(grid, v)
					}
				}
			}

			slices.SortFunc(grid, func(a, b byte) int {
				return int(a) - int(b)
			})

			for i := 0; i < len(grid)-1; i++ {
				if grid[i] == grid[i+1] {
					return false
				}
			}
		}

	}

	return true
}
