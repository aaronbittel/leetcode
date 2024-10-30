package main

func maxVowelsImproved(s string, k int) int {
	max_vowels := 0
	cur_vowels := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if r-l >= k {
			if isVowel(s[l]) {
				cur_vowels--
			}
			l++
		}
		if isVowel(s[r]) {
			cur_vowels++
		}

		if cur_vowels > max_vowels {
			max_vowels = cur_vowels
			// can break here because max_vowels cannot be any greater
			if max_vowels == k {
				break
			}
		}
	}

	return max_vowels
}

func maxVowels(s string, k int) int {
	max_vowels := 0
	cur_vowels := 0
	l := 0
	for r := 0; r < len(s); r++ {
		if r-l >= k {
			if isVowel(s[l]) {
				cur_vowels--
			}
			l++
		}
		if isVowel(s[r]) {
			cur_vowels++
		}

		if cur_vowels > max_vowels {
			max_vowels = cur_vowels
		}
	}

	return max_vowels
}

func isVowel(r byte) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
}
