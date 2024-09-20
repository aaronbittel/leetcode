package main

import "fmt"

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(s, t)
	fmt.Println(isSubsequence(s, t), "expect", true)

	s = "axc"
	t = "ahbgdc"
	fmt.Println(s, t)
	fmt.Println(isSubsequence(s, t), "expect", false)
}
