package main

import (
	"strings"
	"unicode/utf8"
)

func isPalindromeStr(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, utf8.RuneCountInString(s)-1

	valid := func(b byte) bool {
		return (b >= 48 && b <= 57) || (b >= 97 && b <= 122)
	}

	for left <= right {
		cLeft, cRight := s[left], s[right]
		if !valid(cLeft) {
			left++
			continue
		}
		if !valid(cRight) {
			right--
			continue
		}
		if cLeft != cRight {
			return false
		}
		left++
		right--
	}

	return true
}
