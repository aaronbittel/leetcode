package main

func isHappy(n int) bool {
	return isHappyHelper(n, make(map[int]struct{}))
}

func isHappyHelper(n int, visited map[int]struct{}) bool {
	if n == 1 {
		return true
	}

	nextN := 0
	for n > 0 {
		digit := n % 10
		nextN += digit * digit
		n /= 10
	}

	if _, ok := visited[nextN]; ok {
		return false
	}

	visited[nextN] = struct{}{}

	return isHappyHelper(nextN, visited)
}
