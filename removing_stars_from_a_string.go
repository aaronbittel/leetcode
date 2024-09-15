package main

func removeStars(s string) string {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)
}
