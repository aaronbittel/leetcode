package main

import "strings"

func lengthOfLastWord(s string) int {
	count := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if count == 0 {
				continue
			} else {
				break
			}
		} else {
			count++
		}
	}

	return count
}

func lengthOfLastWordSimple(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
