package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	patternToWord := make(map[byte]string)
	wordToPattern := make(map[string]struct{})

	for i, word := range words {
		v, ok := patternToWord[pattern[i]]
		if !ok {
			if _, ok := wordToPattern[word]; ok {
				return false
			}
			wordToPattern[word] = struct{}{}
			patternToWord[pattern[i]] = word
			continue
		}

		if v != word {
			return false
		}
	}

	return true
}
