package main

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	cur := 0
	for i := 0; i < len(t); i++ {
		if s[cur] != t[i] {
			continue
		}
		cur++
		if cur >= len(s) {
			return true
		}
	}
	return false
}
