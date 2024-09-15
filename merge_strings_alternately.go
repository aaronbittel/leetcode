package main

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var b strings.Builder

	for i := 0; i < len(word1) || i < len(word2); i++ {
		if i < len(word1) {
			b.WriteByte(word1[i])
		}
		if i < len(word2) {
			b.WriteByte(word2[i])
		}
	}

	return b.String()
}
