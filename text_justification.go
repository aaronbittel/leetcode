package main

import (
	"math"
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {

	left, right := 0, 0
	curLen := 0

	var result []string

	for right < len(words) {
		if left == right {
			curLen += len(words[left])
			right++
			continue
		}

		if curLen+1+len(words[right]) <= maxWidth {
			curLen += 1 + len(words[right])
			right++
			continue
		}

		result = append(result, otherLine(words[left:right], maxWidth))
		curLen = 0
		left = right
	}

	result = append(result, lastLine(words[left:], maxWidth))

	return result
}

func fullJustifyV1(words []string, maxWidth int) []string {
	lines := make([][]string, 1)

	lines[0] = make([]string, 0)
	lines[0] = append(lines[0], words[0])

	row := 0
	curLen := len(words[0])
	for i := 1; i < len(words); i++ {
		if curLen+1+len(words[i]) <= maxWidth {
			lines[row] = append(lines[row], words[i])
			curLen += 1 + len(words[i])
		} else {
			row++
			lines = append(lines, make([]string, 0))
			lines[row] = append(lines[row], words[i])
			curLen = len(words[i])
		}
	}

	var result []string

	for i, row := range lines {
		if i == len(lines)-1 {
			result = append(result, lastLine(row, maxWidth))
			break
		}
		result = append(result, otherLine(row, maxWidth))
	}

	return result
}

func lastLine(ws []string, maxWidth int) string {
	str := strings.Join(ws, " ")
	str += strings.Repeat(" ", maxWidth-len(str))
	return str
}

func otherLine(ws []string, maxWidth int) string {
	wCount := len(ws)
	spacesCount := maxWidth - len(strings.Join(ws, ""))
	spaces := wCount - 1

	spacesBetween := int(math.Ceil(float64(spacesCount) / float64(spaces)))

	str := ws[0]
	for i := 1; i < wCount; i++ {
		str += strings.Repeat(" ", spacesBetween) + ws[i]
		spacesCount -= spacesBetween
		spaces--
		spacesBetween = int(math.Ceil(float64(spacesCount) / float64(spaces)))
	}

	str += strings.Repeat(" ", maxWidth-len(str))
	return str
}
