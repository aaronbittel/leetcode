package main

func plusOneNewAlloc(digits []int) []int {
	result := make([]int, len(digits))
	copy(result, digits)
	return plusOneInplace(result)
}

func plusOneInplace(digits []int) []int {
	last := len(digits) - 1
	if digits[last] < 9 {
		digits[last] += 1
		return digits
	}

	cur := len(digits) - 1
	for digits[cur] == 9 {
		digits[cur] = 0
		cur -= 1

		if cur < 0 {
			return append([]int{1}, digits...)
		}
	}

	digits[cur] += 1
	return digits
}
