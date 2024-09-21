package main

import "fmt"

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	// fmt.Println(s, numRows)
	// res := convert2(s, numRows)
	// fmt.Println(res, res == "PAHNAPLSIIGYIR")

	s = "PAYPALISHIRING"
	numRows = 4
	fmt.Println(s, numRows)
	res := convert2(s, numRows)
	fmt.Println(res, res == "PINALSIGYAHRPI")

	// s = "ABC"
	// numRows = 2
	// fmt.Println(s, numRows)
	// res = convert2(s, numRows)
	// fmt.Println(res, res == "ACB")
}
