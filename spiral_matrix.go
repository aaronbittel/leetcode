package main

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	left, right := 0, len(matrix[0])
	top, bottom := 0, len(matrix)

	for left < right && top < bottom {
		for i := left; i < right; i++ {
			result = append(result, matrix[top][i])
		}
		top += 1

		for i := top; i < bottom; i++ {
			result = append(result, matrix[i][right-1])
		}
		right -= 1

		if !(left < right && top < bottom) {
			break
		}

		for i := right - 1; i >= left; i-- {
			result = append(result, matrix[bottom-1][i])
		}
		bottom -= 1

		for i := bottom - 1; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left += 1
	}

	return result
}
