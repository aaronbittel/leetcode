package main

import (
	"strings"
)

func minWindow(s string, t string) string {
	res := ""
	countT := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		countT[t[i]]++
	}

	window := make(map[byte]int)
	l := 0
	have, need := 0, len(t)

	for r := 0; r < len(s); r++ {
		c := s[r]
		if !strings.Contains(t, string(c)) {
			continue
		}

		window[c]++
		if window[c] == countT[c] {
			have += window[c]
		}

		for have == need {
			if res == "" || r-l+1 < len(res) {
				res = s[l : r+1]
			}
			if v, ok := countT[s[l]]; ok {
				window[s[l]]--
				if window[s[l]] < v {
					have -= countT[s[l]]
				}
			}
			l++
		}

	}

	return res
}
