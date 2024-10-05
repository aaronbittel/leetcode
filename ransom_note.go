package main

func canConstruct(ransomNote string, magazine string) bool {
	magazineMap := map[rune]int{}
	for _, mag := range magazine {
		magazineMap[mag]++
	}

	for _, ran := range ransomNote {
		if v, ok := magazineMap[ran]; ok && v >= 1 {
			magazineMap[ran]--
		} else {
			return false
		}
	}
	return true
}

