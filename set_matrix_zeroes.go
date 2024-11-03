package main

func setZeroes(matrix [][]int) {
	var (
		rows = make(map[int]struct{})
		cols = make(map[int]struct{})
	)

	for r, row := range matrix {
		for c, el := range row {
			if el != 0 {
				continue
			}

			rows[r] = struct{}{}
			cols[c] = struct{}{}
		}
	}

	for r := range rows {
		for c := 0; c < len(matrix[0]); c++ {
			matrix[r][c] = 0
		}
	}

	for c := range cols {
		for r := 0; r < len(matrix); r++ {
			matrix[r][c] = 0
		}
	}
}
