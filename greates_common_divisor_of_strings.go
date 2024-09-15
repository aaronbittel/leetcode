package main

import (
	"strings"
)

func gcdOfStrings(str1 string, str2 string) string {
	l1, l2 := len(str1), len(str2)
	if str1 == str2 {
		return str1
	}
	candidate := ""
	for i := 0; i < l1 && i < l2; i++ {
		li := i + 1
		c := str1[:i+1]
		if strings.Count(str1, c)*li != l1 {
			continue
		}
		if strings.Count(str2, c)*li != l2 {
			continue
		}
		candidate = c
	}
	return candidate
}
