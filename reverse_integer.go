package main

func reverse(x int) int {
	res := 0
	var digit int
	for x != 0 {
		digit = x % 10
		x /= 10

		res = res*10 + digit
	}

	if res > (1<<31-1) || res < -1<<31 {
		return 0
	}

	return res
}
