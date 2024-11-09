package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	originalLetters := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		originalLetters[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		originalLetters[t[i]]--
	}

	for _, v := range originalLetters {
		if v != 0 {
			return false
		}
	}

	return true
}

func isAnagramOne(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	originalLetters := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		originalLetters[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if v, ok := originalLetters[t[i]]; ok {
			originalLetters[t[i]]--
			if v == 1 {
				delete(originalLetters, t[i])
			}
		}
	}

	return len(originalLetters) == 0
}
