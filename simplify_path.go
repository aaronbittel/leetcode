package main

import (
	"strings"
)

func simplifyPath(path string) string {
	if path == "/" {
		return "/"
	}

	stack := make([]string, 0)

	// ignore starting "/"
	for _, s := range strings.Split(path[1:], "/") {
		switch s {
		case "", ".":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		// dirname
		default:
			stack = append(stack, s)
		}
	}

	return "/" + strings.Join(stack, "/")
}
