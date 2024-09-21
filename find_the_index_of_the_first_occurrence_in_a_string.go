package main

func strStr(haystack string, needle string) int {
outer:
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					continue outer
				}
			}
			return i
		}
	}
	return -1
}
