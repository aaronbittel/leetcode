package main

import (
	"slices"
)

// 1605ms (HashMap) 1371ms (slice-sorting)
func findSubstringWorksButSuperSlow(s string, words []string) []int {
	totalLen := 0
	for _, w := range words {
		totalLen += len(w)
	}

	// slices.Sort(words)

	res := []int{}

	for i := 0; i <= len(s)-totalLen; i++ {
		if check(s[i:i+totalLen], words) {
			res = append(res, i)
		}
	}

	return res
}

func check(s string, words []string) bool {
	wordMap := make(map[string]int)
	for _, w := range words {
		wordMap[w]++
	}

	wLen := len(words[0])
	for i := 0; i <= len(s)-wLen; i += wLen {
		word := s[i : i+wLen]
		if _, ok := wordMap[word]; !ok {
			return false
		}
		wordMap[word]--
	}

	for _, v := range wordMap {
		if v != 0 {
			return false
		}
	}

	return true
}

func checkSorting(s string, words []string) bool {
	ws := []string{}

	wLen := len(words[0])
	for i := 0; i <= len(s)-wLen; i += wLen {
		word := s[i : i+wLen]
		ws = append(ws, word)
	}

	slices.Sort(ws)
	return slices.Equal(ws, words)
}
