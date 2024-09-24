package main

func findRepeatedDnaSequences(s string) []string {
	dnaSet := map[string]int{}
	res := []string{}

	for left := 0; left < len(s)-9; left++ {
		seq := s[left : left+10]
		dnaSet[seq]++
		if c, ok := dnaSet[seq]; ok && c == 2 {
			res = append(res, seq)
		}
	}

	return res
}
