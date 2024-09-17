package main

func longestCommonPrefix(strs []string) string {
	if len(strs[0]) == 0 {
		return ""
	}

	shortest := len(strs[0])
	for _, s := range strs {
		if len(s) == 0 {
			return ""
		}

		if len(s) < shortest {
			shortest = len(s)
		}
	}

	common := ""

	for i := 0; i < shortest; i++ {
		candidate := string(strs[0][i])
		for _, s := range strs {
			if string(s[i]) != candidate {
				return string(common)
			}
		}
		common += candidate
	}

	return string(common)
}
