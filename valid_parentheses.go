package main

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		last := stack[len(stack)-1]
		if last == '(' && c == ')' || last == '[' && c == ']' || last == '{' && c == '}' {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return len(stack) == 0
}
