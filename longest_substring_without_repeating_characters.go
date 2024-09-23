package main

func lengthOfLongestSubstring(s string) int {
	holder := map[rune]int{}
	maxLen := 0
	left := 0
	for right, char := range s {
		if right, ok := holder[char]; ok && right >= left {
			left = right + 1
		}
		holder[char] = right
		currLen := right - left + 1
		if currLen > maxLen {
			maxLen = currLen
		}
	}
	return maxLen
}

// Must faster than map 252 -> 65ms because not that many elements in map
// therefore overhead of creating multiple maps, inserting values, calculating
// hashes is much greater than just loop through a few elements in a slice
func lengthOfLongestSubstringWithSliceInsteadOfMap(s string) int {
	var (
		longest = 0
		cur     int
	)

	for i := 0; i < len(s); i++ {
		cur = 0
		l := make([]byte, 0)
		for j := i; j < len(s); j++ {
			if in(l, s[j]) {
				if cur > longest {
					longest = cur
				}
				break
			} else {
				l = append(l, s[j])
				cur++
			}
		}
		if cur > longest {
			return cur
		}
	}
	return longest
}

func in(slice []byte, b byte) bool {
	for _, s := range slice {
		if s == b {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstringSlowAndLotsOfMemory(s string) int {
	var (
		longest = 0
		cur     int
	)

	for i := 0; i < len(s); i++ {
		cur = 0
		set := make(map[byte]struct{})
		for j := i; j < len(s); j++ {
			if _, ok := set[s[j]]; ok {
				if cur > longest {
					longest = cur
				}
				break
			} else {
				set[s[j]] = struct{}{}
				cur++
			}
		}
		if cur > longest {
			return cur
		}
	}

	return longest
}
