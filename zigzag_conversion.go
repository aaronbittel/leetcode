package main

import (
	"fmt"
	"slices"
)

// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I

func convert2(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	inner := []int{}
	for i := numRows - 1; i < len(s); i += 6 {
		inner = append(inner, i)
	}
	fmt.Println(inner)

	outer := []int{}
	for i := 0; i < len(inner); i++ {
		outer = append(outer, inner[i]-1)
		outer = append(outer, inner[i]+1)
	}
	fmt.Println(outer)

	moreOuter := []int{}
	for i := 0; i < len(outer); i++ {
		moreOuter = append(moreOuter, outer[i]-1)
		moreOuter = append(moreOuter, outer[i]+1)
	}
	fmt.Println(slices.Compact(moreOuter))

	return ""
}

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}
	pattern := make([][]string, numRows)

	for i := range pattern {
		pattern[i] = make([]string, len(s))
	}

	r, c, down := 0, 0, true
	for _, ch := range s {
		pattern[r][c] = string(ch)
		if down {
			r++
			if r == numRows {
				down = false
				r--
			}
		}
		if !down {
			r--
			c++
			if r == 0 {
				down = true
			}
		}
	}

	answer := ""
	for _, p := range pattern {
		for _, c := range p {
			if c == " " {
				continue
			}
			answer += c
		}
	}

	return answer

}
