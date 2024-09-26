package main

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

func romanToInt(s string) int {
	var num int

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			num += 1000
		case 'D':
			num += 500
		case 'C':
			if i+1 >= len(s) {
				num += 100
				break
			}
			if s[i+1] == 'D' {
				num += 400
				i++
			} else if s[i+1] == 'M' {
				num += 900
				i++
			} else {
				num += 100
			}
		case 'L':
			num += 50
		case 'X':
			if i+1 >= len(s) {
				num += 10
				break
			}
			if s[i+1] == 'L' {
				num += 40
				i++
			} else if s[i+1] == 'C' {
				num += 90
				i++
			} else {
				num += 10
			}
		case 'V':
			num += 5
		case 'I':
			if i+1 >= len(s) {
				num += 1
				break
			}
			if s[i+1] == 'V' {
				num += 4
				i++
			} else if s[i+1] == 'X' {
				num += 9
				i++
			} else {
				num += 1
			}
		}
	}

	return num
}
