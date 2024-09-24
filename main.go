package main

import (
	"fmt"
	"slices"
)

func main() {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	expected := []string{"AAAAACCCCC", "CCCCCAAAAA"}
	res := findRepeatedDnaSequences(s)
	fmt.Println(res)
	fmt.Println("res", slices.Compare(expected, res))
	fmt.Println("-------------")

	s = "AAAAAAAAAAAAA"
	expected = []string{"AAAAAAAAAA"}
	res = findRepeatedDnaSequences(s)
	fmt.Println(res)
	fmt.Println("res", slices.Compare(expected, res))
	fmt.Println("-------------")

	s = "AAAAAAAAAAA"
	expected = []string{"AAAAAAAAAA"}
	res = findRepeatedDnaSequences(s)
	fmt.Println(res)
	fmt.Println("res", slices.Compare(expected, res))
	fmt.Println("-------------")

	s = "CAAAAAAAAAC"
	expected = []string{}
	res = findRepeatedDnaSequences(s)
	fmt.Println(res)
	fmt.Println("res", slices.Compare(expected, res))
	fmt.Println("-------------")
}
