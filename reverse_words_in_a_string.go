package main

import (
	"fmt"
	"slices"
	"strings"
)

// 0ms, 3.37mb
func reverseWords(s string) string {
	var res []string

	isSpace := func(b byte) bool {
		return b == ' '
	}

	i := len(s) - 1
	for ; i >= 0 && isSpace(s[i]); i-- {
		// Skip whitespace at the end
	}

	r := i
	for i >= 0 {
		if !isSpace(s[i]) {
			i--
			if i < 0 {
				res = append(res, s[:r+1])
			}
			continue
		}

		// space again
		res = append(res, s[i+1:r+1])

		for i >= 0 && isSpace(s[i]) {
			i--
		}
		r = i
	}

	return strings.Join(res, " ")
}

func reverseWordsV3(s string) string {
	var res string

	isSpace := func(b byte) bool {
		return b == ' '
	}

	i := len(s) - 1
	for ; i >= 0 && isSpace(s[i]); i-- {
		// Skip whitespace at the end
	}

	r := i
	for i >= 0 {
		if !isSpace(s[i]) {
			i--
			if i < 0 {
				res += fmt.Sprintf(" %s", s[:r+1])
			}
			continue
		}

		// space again
		res += fmt.Sprintf(" %s", s[i+1:r+1])

		for i >= 0 && isSpace(s[i]) {
			i--
		}
		r = i
	}

	return res[1:]
}

// 0 ms, 3.22mb
func reverseWordsV2(s string) string {
	words := strings.Fields(s)

	slices.Reverse(words)
	return strings.Join(words, " ")
}

// 2 ms, 3.34mb
func reverseWordsV1(s string) string {
	words := strings.Fields(s)

	var b strings.Builder
	b.WriteString(words[len(words)-1])
	for i := len(words) - 2; i >= 0; i-- {
		b.WriteString(" " + words[i])
	}
	return b.String()
}
