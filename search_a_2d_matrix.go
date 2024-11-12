package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	rows, cols := len(matrix), len(matrix[0])
	l, r := 0, (rows*cols)-1

	for l <= r {
		m := l + (r-l)/2
		row, col := transpone(m, cols)
		fmt.Println(m, row, col)

		cur := matrix[row][col]
		if cur == target {
			return true
		}

		if cur > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return false
}

func transpone(idx, cols int) (row, col int) {
	row = idx / cols
	col = idx % cols
	return
}
